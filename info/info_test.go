package info

import (
	"errors"
	"syscall"
	"strconv"
	"testing"
	"runtime"
)

func TestNumCPUSuccess(t *testing.T) {
	count := strconv.Itoa(runtime.NumCPU())
	result := CPU()

	if result != count {
		t.Errorf("Expected CPU count %s, but got %s", count, result)
	}
}

func TestMemorySuccess(t *testing.T) {
	result, err := Memory(syscall.Sysinfo)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	_, err = strconv.Atoi(result)
	if err != nil {
		t.Errorf("Expected numeric value in memory result, but got %v", result)
	}
}

func TestMemoryError(t *testing.T) {
	mockSysinfo := func(info *syscall.Sysinfo_t) error {
		return errors.New("mock error")
	}

	result, err := Memory(mockSysinfo)
	if err == nil {
		t.Errorf("Expected error due to mock failure, but got none")
	}

	_, err = strconv.Atoi(result)
	if err == nil {
		t.Errorf("Expected non-numeric result due to error, but got numeric value: %v", result)
	}
}
