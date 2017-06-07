package windows

type (
	Bool bool
	Byte byte
	Dword uint32
	DwordPtr uintptr
	DwordLong uint64
	Dword32 int32
	Dword64 int64
	LptStr uint16
	PULong64 uint64
	SizeT uintptr
)

// Handle is the type alias for a standard Windows handle.
type Handle uintptr
