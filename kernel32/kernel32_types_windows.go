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
