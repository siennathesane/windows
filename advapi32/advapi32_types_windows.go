package advapi32

import (
	"github.com/mxplusb/windows"
	"github.com/mxplusb/windows/common"
)

const (
	UNLen = 256
)

// GetUserName retrieves the name of the user associated with the current thread.
func GetUserName() (string, error) {
	var size windows.LpdWord = UNLen + 1
	localBuffer := make([]windows.LptStr, size)
	if err := getUserName(&localBuffer[0], &size); err != nil {
		return "", err
	}
	return common.LptStrToString(localBuffer), nil
}
