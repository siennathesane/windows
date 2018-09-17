package windows

import "syscall"

type (
	Bool bool
	Byte byte
	Dword uint32
	DwordPtr uintptr
	DwordLong uint64
	Dword32 int32
	Dword64 int64
	Handle uintptr
	HpCon uintptr
	LpcStr []uint16
	LpStr uint16
	LptStr uint16
	LpdWord uint32
	LpVoid uintptr
	PULong64 uint64
	Short uint16
	SizeT uintptr
)

func StringToLpcStr(s string) ([]uint16, error) {
	return syscall.UTF16FromString(s)
}

type COORD struct {
	X Short
	Y Short
}

// ProcessInformation aligns with PROCESS_INFORMATION. See: https://docs.microsoft.com/en-us/windows/desktop/api/processthreadsapi/ns-processthreadsapi-_process_information
type ProcessInformation struct {
	HProcess   Handle
	HThread    Handle
	DwProcess  Dword
	DwThreadId Dword
}

type SecurityAttributes struct {
	NLength              Dword
	LpSecurityDescriptor uintptr
	BInheritHandle       bool
}

type StartupInfoA struct {
	Cb         Dword
	LpReserved LpcStr
	LpDesktop LpStr
	LpTitle LpStr
	DwX Dword
	
	LPSTR      lpDesktop;
	LPSTR      lpTitle;
	DWORD      dwX;
	DWORD      dwY;
	DWORD      dwXSize;
	DWORD      dwYSize;
	DWORD      dwXCountChars;
	DWORD      dwYCountChars;
	DWORD      dwFillAttribute;
	DWORD      dwFlags;
	WORD       wShowWindow;
	WORD       cbReserved2;
	LPBYTE     lpReserved2;
	HANDLE     hStdInput;
	HANDLE     hStdOutput;
	HANDLE     hStdError;
}
STARTUPINFOA
