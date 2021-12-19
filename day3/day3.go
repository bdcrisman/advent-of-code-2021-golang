package day3

import (
	"log"
	"strconv"

	"github.com/bdcrisman/advent-of-code/2021/utils"
)

func Day3(inputPath string, isPart1 bool) int64 {
	data, err := utils.ReadAsTwoDimStrArray(inputPath)
	if err != nil {
		log.Fatal(err)
		return -1
	}

	gamma, epsilon := findGammaEpsilon(data, isPart1)
	return gamma * epsilon
}

func findGammaEpsilon(data [][]string, isPart1 bool) (int64, int64) {
	var gamma string
	var epsilon string

	numCols := len(data[0])

	for col := 0; col < numCols; col++ {
		var zeroes int
		var ones int
		var next int

		for i := 0; i < len(data); i++ {
			v, err := strconv.Atoi(data[i][col])
			if err != nil {
				continue
			}

			if v > 0 {
				ones++
			} else {
				zeroes++
			}
		}

		if ones > zeroes {
			next = 1
		} else if zeroes > ones {
			next = 0
		}

		gamma += strconv.Itoa(next)
		if next > 0 {
			epsilon += strconv.Itoa(0)
		} else {
			epsilon += strconv.Itoa(1)
		}
	}

	g, _ := strconv.ParseInt(gamma, 2, 64)
	e, _ := strconv.ParseInt(epsilon, 2, 32)

	return g, e
}
