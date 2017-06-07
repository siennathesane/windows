package kernel32

import (
	"fmt"
	"testing"

	"github.com/mxplusb/windows"
	"github.com/mxplusb/windows/common"
	"syscall"
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
	t.Logf("run for %d cycles thus far\n", cycles)
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

func TestInternalGetComputerName(t *testing.T) {
	var size windows.LpdWord = MaxComputerNameLength + 1
	localBuffer := make([]windows.LptStr, size)
	if err := getComputerName(&localBuffer[0], &size); err != nil {
		t.Fatal(err)
	}
	// just in case.
	if empty := string(common.LptStrToString(localBuffer[0:size])); empty == "" {
		t.Fatal("computer name is empty!")
	}
	t.Logf("computer name is %s", string(common.LptStrToString(localBuffer[:size])))
}

func TestGetComputerName(t *testing.T) {
	me, err := GetComputerName()
	if err != nil {
		t.Fatal(err)
	}
	if me == "" {
		t.Fatal("computer name is empty!")
	}
	t.Logf("computer name is %s", me)
}

func TestLoadLibrary(t *testing.T) {
	lpHandle, err := LoadLibrary("psapi.dll")
	if err != nil {
		t.Fatal(err)
	}
	if lpHandle == 0 {
		t.Fatal("handle is empty")
	}
	t.Logf("handle for psapi.dll is %d", lpHandle)
}

func TestGetProcAddress(t *testing.T) {
	// load something else not already loaded.
	psapi, err := LoadLibrary("psapi.dll")
	if err != nil {
		t.Fatal(err)
	}

	var byteRef *byte
	byteRef, err = syscall.BytePtrFromString("GetPerformanceInfo")
	if err != nil {
		t.Fatal(err)
	}

	addr, err := GetProcAddress(psapi, byteRef)
	if err != nil {
		t.Fatal(err)
	}
	if addr == 0 {
		t.Fatal("proc address is 0")
	}
	t.Logf("GetperformanceInfo is at %d", addr)
}

func ExampleGetProcAddress() {
	// let's load a DLL to see if there's a function in it.
	psapiDll, err := LoadLibrary("psapi.dll")
	if err != nil {
		panic(err)
	}

	// because GetProcAddress requires a byte pointer, let's do the conversion.
	// we're looking for the GetPerformanceInfo API.
	var byteRef *byte
	byteRef, err = syscall.BytePtrFromString("GetPerformanceInfo")
	if err != nil {
		panic(err)
	}

	addr, err := GetProcAddress(psapiDll, byteRef)
	if err != nil {
		panic(err)
	}
	fmt.Printf("GetPerformanceInfo is at %d", addr)
}

func TestGetVersion(t *testing.T) {
	ver, err := GetVersion()
	if err != nil {
		t.Fatal(err)
	}
	if ver == 0 {
		t.Fatal("version is 0")
	}
	t.Logf("tested on Windows %d", ver)
}
