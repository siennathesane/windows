package windows

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
	LpByte byte
	LpDword uint32
	LpcStr uint16
	LpStr uint16
	LptStr uint16
	LpdWord uint32
	LpVoid uint16
	LpVoidByte byte
	PULong64 uint64
	Short uint16
	SizeT uintptr
	Word uint16
)

type COORD struct {
	X Short
	Y Short
}

// Overlapped contains information used in asynchronous (or overlapped) input and output (I/O). See: https://docs.microsoft.com/en-us/windows/desktop/api/minwinbase/ns-minwinbase-_overlapped
type Overlapped struct {
	Internal     uintptr
	InternalHigh uintptr
	Offset       uint32
	OffsetHigh   uint32
	HEvent       Handle
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

type StartupInfo struct {
	Cb              Dword
	LpReserved      LpcStr
	LpDesktop       LpStr
	LpTitle         LpStr
	DwX             Dword
	DwY             Dword
	DwXSize         Dword
	DwYSize         Dword
	DwXCountChars   Dword
	DwYCountChars   Dword
	DwFillAttribute Dword
	DwFlags         Dword
	WShowWindow     Word
	CbReserved2     Word
	LpReserved2     LpByte
	HStdInput       Handle
	HStdOutput      Handle
	HStdError       Handle
}
