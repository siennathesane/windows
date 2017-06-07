// MACHINE GENERATED BY 'go generate' COMMAND; DO NOT EDIT

package kernel32

import (
	"syscall"
	"unsafe"

	"github.com/mxplusb/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return nil
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	modkernel32 = windows.NewLazySystemDLL("kernel32.dll")

	procGetCurrentProcess     = modkernel32.NewProc("GetCurrentProcess")
	procQueryProcessCycleTime = modkernel32.NewProc("QueryProcessCycleTime")
)

// GetCurrentProcess retrieves a pseudo handle for the current process. See: https://msdn.microsoft.com/en-us/library/windows/desktop/ms683179%28v=vs.85%29.aspx
func GetCurrentProcess() (pseudoHandle windows.HANDLE, err error) {
	r0, _, e1 := syscall.Syscall(procGetCurrentProcess.Addr(), 0, 0, 0, 0)
	pseudoHandle = windows.HANDLE(r0)
	if pseudoHandle == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

// QueryProcessCycleTime retrieves the sum of the cycle time of all threads of the specified process.
func QueryProcessCycleTime(handle windows.HANDLE, cycleTime *windows.PULONG64) (err error) {
	r1, _, e1 := syscall.Syscall(procQueryProcessCycleTime.Addr(), 2, uintptr(handle), uintptr(unsafe.Pointer(cycleTime)), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}