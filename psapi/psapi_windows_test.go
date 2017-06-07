package psapi

import (
	"fmt"
	"testing"

	"github.com/mxplusb/windows/kernel32"
)

func TestGetProcessMemoryInfo(t *testing.T) {
	localMemCounter := NewProcessMemoryCounter()
	me, err := kernel32.GetCurrentProcess()
	if err != nil {
		t.Fatal(err)
	}
	if err := getProcessMemoryInfo(me, localMemCounter, localMemCounter.Cb); err != nil || localMemCounter.WorkingSetSize == 0 {
		t.Fatal(err)
	}
	t.Logf("using ~%d bytes of memory", localMemCounter.WorkingSetSize)
}

func TestProcessMemoryCounters_GetProcessMemoryInfo(t *testing.T) {
	localMemCounter := NewProcessMemoryCounter()
	me, err := kernel32.GetCurrentProcess()
	if err != nil {
		t.Fatal(err)
	}
	if err := localMemCounter.GetProcessMemoryInfo(me); err != nil || localMemCounter.WorkingSetSize == 0 {
		t.Fatal(err)
	}
	t.Logf("using ~%d bytes of memory", localMemCounter.WorkingSetSize)
}

func ExampleGetProcessMemoryInfo() {
	localMemoryCounter := NewProcessMemoryCounter()
	me, err := kernel32.GetCurrentProcess()
	if err != nil {
		panic(err)
	}
	if err := localMemoryCounter.GetProcessMemoryInfo(me); err != nil {
		panic(err)
	}
	fmt.Printf("using %d as the working set!", localMemoryCounter.WorkingSetSize)
}

func TestGetPerformanceInfo(t *testing.T) {
	localPerfInfo := NewPerformanceInformation()
	if err := getPerformanceInfo(localPerfInfo, localPerfInfo.Cb); err != nil || localPerfInfo.HandleCount == 0 {
		t.Fatal(err)
	}
	t.Logf("%d handles are seen in the system", localPerfInfo.HandleCount)
}

func TestPerformanceInformation_GetPerformanceInformation(t *testing.T) {
	localPerfInfo := NewPerformanceInformation()
	if err := getPerformanceInfo(localPerfInfo, localPerfInfo.Cb); err != nil {
		t.Fatal(err)
	}
	if err := localPerfInfo.GetPerformanceInformation(); err != nil || localPerfInfo.HandleCount == 0 {
		t.Fatal(err)
	}
	t.Logf("%d handles are seen in the system", localPerfInfo.HandleCount)
}
