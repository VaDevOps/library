package info

import (
	"strconv"
	"testing"
)

func TestMemory(t *testing.T) {
	result, err := Memory()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if _, err := strconv.Atoi(result); err != nil {
		t.Errorf("Expected numeric value, but got %v", result)
	}
}

func TestCPU() {

}
