package kernel32

import (
	"errors"
	"github.com/mxplusb/windows"
	"github.com/mxplusb/windows/common"
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

type BackgroundProcess struct {
	readPipeHandle     windows.Handle
	writePipeHandle    windows.Handle
	procInfo           windows.ProcessInformation
	startupInfo        windows.StartupInfo
	BufferSize         windows.Dword
	ChildProcessHandle windows.Handle
}

func (bg *BackgroundProcess) CreateBackgroundProcess(lpCommandLine string) (error) {
	// build some necessities.
	var lpProcessAttrs windows.SecurityAttributes
	//lpProcessAttrs.BInheritHandle = true
	var lpThreadAttrs windows.SecurityAttributes

	lpCli, err := common.StringToLpStr(lpCommandLine)
	if err != nil {
		return err
	}

	// create the pipe for the process to be read.
	var lpPipeSecAttrs windows.SecurityAttributes
	CreatePipe(&bg.readPipeHandle, &bg.writePipeHandle, &lpPipeSecAttrs, bg.BufferSize)

	if ok := CreateProcess(nil, lpCli, &lpProcessAttrs, &lpThreadAttrs, true, 0, nil, nil, &bg.startupInfo, &bg.procInfo); !ok {
		return GetLastError()
	}

	// this is where we save the child process handle.
	if bg.procInfo.HProcess <= 0 {
		return errors.New("child handle does not exist")
	}
	return nil
}

// GetOutput will get all the output from the child process created.
//func (bg *BackgroundProcess) GetOutput(outBuf []byte, readSize windows.Dword) error {
//	var overlapped windows.Overlapped
//	if readSize > bg.BufferSize {
//		return errors.New("cannot read outsize buffer size")
//	}
//	if ok := ReadFile(bg., outBuf, readSize, &overlapped); !ok {
//		return GetLastError()
//	}
//	return nil
//}

func (bg *BackgroundProcess) CloseBackgroundProcess() error {
	bg.ChildProcessHandle = OpenProcess(DeleteProcessAccessRight, false, bg.procInfo.DwProcess)
	if bg.ChildProcessHandle <= 0 {
		return GetLastError()
	}
	if ok := TerminateProcess(bg.ChildProcessHandle, 0); !ok {
		return GetLastError()
	}
	return nil
}
