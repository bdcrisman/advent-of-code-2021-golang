package day1

import (
	"path/filepath"

	"github.com/bdcrisman/advent-of-code/2021/utils"
)

// Gets the number of increases occurring in the data.
func Day1(inputPath string, isPart1 bool) int {
	path := filepath.Join("day1", "")
	data, err := utils.ReadAsIntArray(path)
	if err != nil {
		println(err.Error())
		return 0
	}

	if isPart1 {
		return countIncreasesPart1(data)
	}
	return countIncreasesPart2(data)
}

// Counts the number of increases for part 1.
func countIncreasesPart1(data []int) int {
	if len(data) <= 0 {
		return 0
	}

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

// Counts the number of increases for part 2.
func countIncreasesPart2(data []int) int {
	if len(data) < 3 {
		return 0
	}

	var count int
	last := data[0] + data[1] + data[2]
	lastIndex := len(data) - 1

	for i := 0; i < len(data); i++ {
		if i+2 > lastIndex {
			break
		}

		sum := data[i] + data[i+1] + data[i+2]
		if sum > last {
			count++
		}
		last = sum
	}

	return count
}
