package day1

import (
	"path/filepath"

	"github.com/bdcrisman/advent-of-code/2021/utils"
)

func Day1(inputPath string) []int {
	if inputPath == "" {
		println("DAY1: no input path given")
		return nil
	}

	path := filepath.Join("day1", inputPath)
	data, err := utils.ReadAsIntArray(path)
	if err != nil {
		println(err.Error())
		return nil
	}

	return data
}
