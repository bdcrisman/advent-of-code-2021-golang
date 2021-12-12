package main

import (
	"testing"

	"github.com/bdcrisman/advent-of-code/2021/utils"
)

func TestFileScannerPass(t *testing.T) {
	scanner, _ := utils.FileScanner("day1/example")
	if scanner == nil {
		t.Log("should not be nil")
	}
	t.Log("pass")
}
