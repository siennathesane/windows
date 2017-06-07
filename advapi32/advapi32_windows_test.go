package advapi32

import (
	"testing"
	"github.com/mxplusb/windows/common"
	"github.com/mxplusb/windows"
)

func TestInternalGetUserName(t *testing.T) {
	var size windows.LpdWord = UNLen + 1
	localBuffer := make([]windows.LptStr, size)
	if err := getUserName(&localBuffer[0], &size); err != nil {
		t.Fatal(err)
	}
	// just in case.
	if empty := string(common.LptStrToString(localBuffer[0:size])); empty == "" {
		t.Fatal("user name is empty!")
	}
	t.Logf("user name is %s", string(common.LptStrToString(localBuffer[:size])))
}

func TestGetUserName(t *testing.T) {
	me, err := GetUserName()
	if err != nil {
		t.Fatal(err)
	}
	if me == "" {
		t.Fatal("user name is empty")
	}
	t.Logf("user name is %s", me)
}