package psapi

import (
	"github.com/mxplusb/windows"
	"github.com/mxplusb/windows/common"
)

// ProcessMemoryCounters contains the memory statistics for a process.
type ProcessMemoryCounters struct {
	// The size of the structure, in bytes.
	Cb windows.Dword
	// The number of page faults.
	PageFaultCount windows.Dword
	// The peak working set size, in bytes.
	WorkingSetSize windows.SizeT
	// The current working set size, in bytes.
	QuotaPeakPagedPoolUsage windows.SizeT
	// The peak paged pool usage, in bytes.
	QuotaPagedPoolUsage windows.SizeT
	// The current paged pool usage, in bytes.
	QuotaPeakNonPagedPoolUsage windows.SizeT
	// The peak nonpaged pool usage, in bytes.
	QuotaNonPagedPoolUsage windows.SizeT
	// The Commit Charge value in bytes for this process. Commit Charge is the total amount of memory that the memory manager has committed for a running process.
	PagefileUsage windows.SizeT
	// The peak value in bytes of the Commit Charge during the lifetime of this process.
	PeakPagefileUsage windows.SizeT
}

// NewProcessMemoryCounter returns a ProcessMemoryCounter with the appropriate struct size predefined.
func NewProcessMemoryCounter() *ProcessMemoryCounters {
	// it's really interesting, `unsafe.Sizeof` will report this as 8 bytes, but Windows doesn't like it if I don't
	// explicitly set at least 128. weird idiosyncrasy.
	return &ProcessMemoryCounters{Cb: 128}
}

// GetProcessMemoryInfo retrieves information about the memory usage of the specified process.
func (p *ProcessMemoryCounters) GetProcessMemoryInfo(h windows.Handle) error {
	localProcCount := NewProcessMemoryCounter()
	if err := getProcessMemoryInfo(h, localProcCount, localProcCount.Cb); err != nil {
		return err
	}
	if err := common.DeepCopy(p, localProcCount); err != nil {
		return err
	}
	return nil
}

// PerformanceInformation contains performancec information regarding the overall resource commitment to a machine.
type PerformanceInformation struct {
	// The size of this structure, in bytes.
	Cb windows.Dword
	// The number of pages currently committed by the system. Note that committing
	// pages (using VirtualAlloc with MEM_COMMIT) changes this value immediately;
	// however, the physical memory is not charged until the pages are accessed.
	CommitTotal windows.SizeT
	// The current maximum number of pages that can be committed by the system
	// without extending the paging file(s). This number can change if memory is added
	// or deleted, or if pagefiles have grown, shrunk, or been added. If the paging file can
	// be extended, this is a soft limit.
	CommitLimit windows.SizeT
	// The maximum number of pages that were simultaneously in the committed state since the last system reboot.
	CommitPeak windows.SizeT
	// The amount of actual physical memory, in pages.
	PhysicalTotal windows.SizeT
	// The amount of physical memory currently available, in pages. This is the amount of
	// physical memory that can be immediately reused without having to write its contents
	// to disk first. It is the sum of the size of the standby, free, and zero lists.
	PhysicalAvailable windows.SizeT
	// The amount of system cache memory, in pages. This is the size of the standby list plus the system working set.
	SystemCache windows.SizeT
	// The sum of the memory currently in the paged and nonpaged kernel pools, in pages.
	KernelTotal windows.SizeT
	// The memory currently in the paged kernel pool, in pages.
	KernelPaged windows.SizeT
	// The memory currently in the nonpaged kernel pool, in pages.
	KernelNonpaged windows.SizeT
	// The size of a page, in bytes.
	PageSize windows.SizeT
	// The current number of open handles.
	HandleCount windows.Dword
	// The current number of processes.
	ProcessCount windows.Dword
	// The current number of threads.
	ThreadCount windows.Dword
}

// NewPerformanceInfo generates a new struct with the appropriate memory values.
func NewPerformanceInformation() *PerformanceInformation {
	return &PerformanceInformation{Cb: 128}
}

func (p *PerformanceInformation) GetPerformanceInformation() error {
	localPerfInfo := NewPerformanceInformation()
	if err := getPerformanceInfo(localPerfInfo, localPerfInfo.Cb); err != nil {
		return err
	}
	if err := common.DeepCopy(p, localPerfInfo); err != nil {
		return err
	}
	return nil
}

// EnumPageFileInformation contains information about a pagefile
type EnumPageFileInformation struct {
	// The size of this structure, in bytes.
	Cb windows.Dword
	// This member is reserved.
	Reserved windows.Dword
	// The total size of the pagefile, in pages.
	TotalSize windows.SizeT
	// The current pagefile usage, in pages.
	TotalInUse windows.SizeT
	// The peak pagefile usage, in pages.
	PeakUsage windows.SizeT
}

