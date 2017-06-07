package kernel32

import (
	"fmt"
	"testing"

	"github.com/mxplusb/windows"
)

func TestGetCurrentProcess(t *testing.T) {
	pHandle, err := GetCurrentProcess()
	if err != nil || pHandle == 0 {
		t.Fatal(err)
	}
	t.Logf("current runtime pseudohandle is %d", pHandle)
}

func ExampleGetCurrentProcess() {
	pHandle, err := GetCurrentProcess()
	if err != nil {
		panic(err)
	}
	fmt.Println(pHandle)
}

func TestQueryProcessCycleTime(t *testing.T) {
	var cycles windows.PULong64
	me, err := GetCurrentProcess()
	if err != nil {
		t.Fatal(err)
	}
	if err := QueryProcessCycleTime(me, &cycles); err != nil {
		if cycles == 0 {
			t.Fatalf("didn't report any cycles! %s", err)
		}
	}
	t.Logf("ran for %d cycles!\n", cycles)
}

func ExampleQueryProcessCycleTime() {
	var cycles windows.PULong64
	pseudoHandle, err := GetCurrentProcess()
	if err != nil {
		panic(err)
	}
	if err := QueryProcessCycleTime(pseudoHandle, &cycles); err != nil {
		panic(err)
	}
	fmt.Printf("the runtime is using %d cycles\n", cycles)
}