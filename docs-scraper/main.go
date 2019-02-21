package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"sync"
)

const (
	windowsDesktopApiBase = "https://docs.microsoft.com/en-us/windows/desktop/api/index"
)

var (
	// our various matchers.
	remarksRegex = regexp.MustCompile("[A-Z0-9_()]+")
	dllRegex     = regexp.MustCompile("[A-Za-z].*\\.dll")
	headerRegex  = regexp.MustCompile("[A-Za-z].*\\.h")
	libRegex     = regexp.MustCompile("[A-Za-z].*\\.lib")

	// global separator
	globalSep = "^"
)

type FuncExpr struct {
	Code           string   `json:"code"`
	Type           string   `json:"type"`
	DLL            string   `json:"dll"`
	DesktopVersion string   `json:"desktop_version"`
	ServerVersion  string   `json:"server_version"`
	Header         string   `json:"header"`
	Documentation  string   `json:"documentation"`
	Lib            string   `json:"lib"`
	Feature        string   `json:"feature"`
	Remarks        []Remark `json:"remarks"`
}

type Remark struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func main() {
	// set up our pipeline
	indexSprawlerToDesktopApiSetIndices := make(chan string, 100) // results from indexSprawler
	apiSetIndexToFnExtractor := make(chan string, 100)            // tree from the apiSetIndexCrawler
	fnChan := make(chan *FuncExpr, 25)                            // final receiver.

	allTheCode := make([]*FuncExpr, 0) // final code list.

	var wg sync.WaitGroup

	// our main background loop for the channel pipelining.
	go func() {
		for apiSetIndex := range indexSprawlerToDesktopApiSetIndices {
			wg.Add(1)
			go apiSetIndexCrawler("desktop", apiSetIndex, &wg, apiSetIndexToFnExtractor)
		}
	}()

	go func() {
		for fnExtractor := range apiSetIndexToFnExtractor {
			wg.Add(1)
			go codeExtractor(fnExtractor, &wg, fnChan)
		}
	}()

	// TODO (mxplusb): this should stream more efficiently.
	go func() {
		wg.Add(1)
		for fn := range fnChan {
			allTheCode = append(allTheCode, fn)
			functions, err := json.Marshal(allTheCode)
			if err != nil {
				panic(err)
			}
			if err := ioutil.WriteFile("functions.json", functions, 0644); err != nil {
				panic(err)
			}
		}
		wg.Done()
	}()

	wg.Add(1)
	go indexSprawler("desktop", windowsDesktopApiBase, &wg, indexSprawlerToDesktopApiSetIndices)
	wg.Wait()
}

// starts at the Server or Desktop index and gets all known APIs.
func indexSprawler(platform, url string, wg *sync.WaitGroup, out chan string) error {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	z := html.NewTokenizer(resp.Body)
	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// assume end of doc.
			return nil
		case tt == html.StartTagToken:
			t := z.Token()
			isHRef := t.Data == "a"
			if isHRef && len(t.Attr) == 2 {
				if t.Attr[1].Val != "relative-path" {
					continue
				}
				out <- fmt.Sprintf("https://docs.microsoft.com/en-us/windows/%s/api/%s", platform, t.Attr[0].Val)
			}
		}
	}
}

func apiSetIndexCrawler(platform, url string, wg *sync.WaitGroup, out chan string) error {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	z := html.NewTokenizer(resp.Body)
	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// assume end of doc.
			return nil
		case tt == html.StartTagToken:
			t := z.Token()
			isHRef := t.Data == "a"
			if isHRef && len(t.Attr) == 2 {
				// we only care about APIs.
				if t.Attr[1].Val != "relative-path" {
					continue
				}
				// no need to walk up the tree.
				if strings.Contains(t.Attr[0].Val, "index") {
					continue
				}
				var localArtifactIndex string
				if strings.Contains(t.Attr[0].Val, "..") {
					localArtifactIndex = strings.TrimLeft(t.Attr[0].Val, "../")
				} else {
					localArtifactIndex = t.Attr[0].Val
				}
				out <- fmt.Sprintf("https://docs.microsoft.com/en-us/windows/%s/api/%s", platform, localArtifactIndex)
			}
		}
	}
}

// extract the code snippets and the table.
func codeExtractor(url string, wg *sync.WaitGroup, out chan *FuncExpr) error {
	defer wg.Done()
	fn := &FuncExpr{}
	fn.Documentation = url

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	var codeChunk func(*html.Node)
	var buf bytes.Buffer
	codeChunk = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "code" {
			buf.WriteString(n.FirstChild.Data + globalSep)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			codeChunk(c)
		}
	}
	codeChunk(doc)

	// split out the function signature and flush the buffer.
	codeSignature := buf.String()
	codeSnippetSplit := strings.Split(codeSignature, globalSep)
	fn.Code = codeSnippetSplit[0]
	buf.Reset()

	var versionTable func(*html.Node)
	versionTable = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "td" {
			// write each found table value and our custom string separator.
			buf.WriteString(n.FirstChild.Data + globalSep)
		}
		if n.Type == html.ElementNode {
			versionTable(n.FirstChild)
		}
		if n.NextSibling != nil {
			versionTable(n.NextSibling)
		}
	}
	versionTable(doc)

	fn.remarker(buf.String())
	out <- fn
	fmt.Printf("%#v\n", fn)
	return nil
}

// grab the dll, header, lib, and any other remarks.
func (fn *FuncExpr) remarker(s string) {
	splitter := strings.Split(s, globalSep)
	// TODO (mxplusb): figure out why this only works sometimes and not others. it seems to be specific to the DX libraries for some reason.
	for idx := range splitter {
		switch {
		case dllRegex.MatchString(splitter[idx]):
			fn.DLL = splitter[idx]
		case libRegex.MatchString(splitter[idx]):
			fn.Lib = splitter[idx]
		case headerRegex.MatchString(splitter[idx]):
			fn.Header = splitter[idx]
		case remarksRegex.MatchString(splitter[idx]):
			fn.Remarks = append(fn.Remarks, Remark{splitter[idx], splitter[idx+1]})
		}
	}
}
