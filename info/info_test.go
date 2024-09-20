package info

import (
	"errors"
	"syscall"
	"strconv"
	"testing"
	"runtime"
)

func TestNumCPUSuccess(t *testing.T){
	count := strconv.Itoa(runtime.NumCPU())
	result := CPU()

	if result != count {
		t.Errorf("result %s != count %s",result,count)
	}
}

func TestMemorySuccess(t *testing.T) {
	result, err := Memory(syscall.Sysinfo)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	_, err = strconv.Atoi(result) 
	if err != nil {
		t.Errorf("Expected numeric value, but got %v", result)
	}
}

func TestMemoryError(t *testing.T){
	var mockSysinfo = func(info *syscall.Sysinfo_t) error {
		return errors.New("mock error")
	}

	result,err := Memory(mockSysinfo)
	if err == nil {
		t.Errorf("Error is: %v",err)
	}

	_,err = strconv.Atoi(result)
	if err == nil {
		t.Errorf("Result is: %v",result)
	}
}
