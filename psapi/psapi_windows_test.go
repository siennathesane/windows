package psapi

import (
	"testing"
	"unsafe"

	"github.com/mxplusb/windows"
	"github.com/mxplusb/windows/kernel32"
)

func TestGetProcessMemoryInfo(t *testing.T) {
	localMemCounter := new(ProcessMemoryCounters)
	me, err := kernel32.GetCurrentProcess()
	if err != nil {
		t.Fatal(err)
	}
	if err := GetProcessMemoryInfo(me, localMemCounter, windows.DWORD_PTR(unsafe.Sizeof(ProcessMemoryCounters{}))); err != nil || localMemCounter.WorkingSetSize == 0 {
		t.Fatal(err)
	}
}
