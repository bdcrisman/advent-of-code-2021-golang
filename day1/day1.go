package day1

import (
	"path/filepath"

	"github.com/bdcrisman/advent-of-code/2021/utils"
)

// Gets the number of increases occurring in the data.
func Day1Part1(inputPath string) int {
	if inputPath == "" {
		println("DAY1: no input path given")
		return 0
	}

	path := filepath.Join("day1", inputPath)
	data, err := utils.ReadAsIntArray(path)
	if err != nil {
		println(err.Error())
		return 0
	}

	return countNumIncreases(data)
}

func countNumIncreases(data []int) int {
	var count int
	last := data[0]

	for _, i := range data {
		if i > last {
			count++
		}
		last = i
	}

	return count
}
