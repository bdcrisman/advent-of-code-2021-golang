package day3

import (
	"strconv"

	"github.com/bdcrisman/advent-of-code/2021/utils"
)

func Day3(inputPath string, isPart1 bool) int {
	data, err := utils.ReadAsTwoDimIntArray(inputPath)
	if err != nil {
		return -1
	}

	gammaBin, epsilonBin := findGammaEpsilon(data, isPart1)
	gamma, err := strconv.

	return gamma * epsilon
}

func findGammaEpsilon(data [][]int, isPart1 bool) (int, int) {
	var zeroCount int
	var oneCount int
	var gamma int
	var epsilon int

	for i := 0; i < len(data); i++ {
		var zeroes int
		var ones int

		for j := 0; j < len(data[i]); j++ {

		}

		if zeroes > ones {
			zeroCount++
		} else if ones > zeroes {
			oneCount++
		} else {
			println("zeroes and ones are equal?")
		}
	}

	strconv.ParseInt("10110", 2, 64)

	return
}

func calculateGammaEpsilon(data []int, index int, ispart1 bool) int {

}
