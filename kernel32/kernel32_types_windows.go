package kernel32

import (
	"github.com/mxplusb/windows"
	"github.com/mxplusb/windows/common"
)

const (
	MaxComputerNameLength = 31
)

// GetComputerName retrieves the NetBIOS name of the local computer. This name is established at system startup, when the system reads it from the registry. See: https://msdn.microsoft.com/en-us/library/windows/desktop/ms724295(v=vs.85).aspx
func GetComputerName() (string, error) {
	var size windows.LpdWord = MaxComputerNameLength + 1
	localBuffer := make([]windows.LptStr, size)
	if err := getComputerName(&localBuffer[0], &size); err != nil {
		return "", err
	}
	return common.LptStrToString(localBuffer), nil
}

type SecurityAttributes struct {
	// The size, in bytes, of this structure.
	nLength              windows.Dword
	// A pointer to a SECURITY_DESCRIPTOR structure that controls access to the object. If the value of this member is NULL, the object is assigned the default security descriptor associated with the access token of the calling process.
	lpSecurityDescriptor windows.LpVoid
	// A Boolean value that specifies whether the returned handle is inherited when a new process is created.
	bInheritHandle       bool
}

func NewSecurityAttributes() *SecurityAttributes {
	return &SecurityAttributes{nLength: 64}
}