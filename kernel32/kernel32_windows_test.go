package kernel32

import (
	"testing"
	"time"
)

func TestBackGroundProcessAll(t *testing.T) {
	testBg := &BackgroundProcess{
		BufferSize: 1024,
	}

	err := testBg.CreateBackgroundProcess("C:\\Windows\\System32\\calc.exe")
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	time.Sleep(time.Second * 1)

	err = testBg.CloseBackgroundProcess()
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
