package day3

import (
	"log"
	"sort"
	"strconv"

	"github.com/bdcrisman/advent-of-code/2021/utils"
)

func Day3(inputPath string, isPart1 bool) int64 {
	if isPart1 {
		return part1(inputPath)
	}
	return part2(inputPath)
}

func part1(path string) int64 {
	data, err := utils.ReadAsTwoDimStrArray(path)
	if err != nil {
		log.Fatal(err)
		return -1
	}
	gamma, epsilon := findGammaEpsilon(data)
	return gamma * epsilon
}

func findGammaEpsilon(data [][]string) (int64, int64) {
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

func part2(path string) int64 {
	data, err := utils.ReadAsStringArray(path)
	if err != nil {
		return -1
	}
	return findLifeSupportRating(data)
}

func findLifeSupportRating(data []string) int64 {
	var o2list []string
	var co2list []string
	var zeroCount int

	sort.Strings(data)
	for i, row := range data {
		v, _ := strconv.Atoi(string(row[0]))
		if v == 0 {
			zeroCount++
			continue
		}

		if zeroCount > len(data)/2 {
			o2list = make([]string, i)
			co2list = make([]string, len(data)-i)
			copy(o2list, data[:i])
			copy(co2list, data[i:])
		} else {
			o2list = make([]string, len(data)-i)
			co2list = make([]string, i)
			copy(o2list, data[i:])
			copy(co2list, data[:i])
		}
		break
	}

	o2 := findRating(o2list, 1, true)
	co2 := findRating(co2list, 1, false)
	return o2 * co2
}

func findRating(data []string, index int, isO2Rating bool) int64 {
	if len(data) == 1 {
		println("final:", data[0])
		v, _ := strconv.ParseInt(data[0], 2, 64)
		return v
	}

	var zeroCount int
	var lastIndex int

	sort.Strings(data)
	for i, row := range data {
		println(index, ":", row)
		v, _ := strconv.Atoi(string(row[index]))
		if v == 0 {
			zeroCount++
			continue
		}
		lastIndex = i
		break
	}

	var temp []string

	// somehow check if there are 2 values left and if has both 1 and 0...
	// if o2, select 1, else select 0
	half := len(data) / 2

	if zeroCount > half {
		temp = make([]string, lastIndex)
		copy(temp, data[:lastIndex])
	} else if zeroCount == half {

	} else {
		temp = make([]string, len(data)-lastIndex)
		copy(temp, data[lastIndex:])
	}

	return findRating(temp, index+1, isO2Rating)
}
