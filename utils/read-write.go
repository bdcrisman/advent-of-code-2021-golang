package utils

import (
	"bufio"
	"io/ioutil"
	"strconv"
	"strings"
)

// Reads data from file in as an int array.
func ReadAsIntArray(path string) ([]int, error) {
	scanner, err := FileScanner(path)
	if err != nil {
		return nil, err
	}

	var result []int

	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}

	return result, scanner.Err()
}

// Reads in file and returns the scanner.
func FileScanner(path string) (*bufio.Scanner, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	r := strings.NewReader(string(data))
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	return scanner, nil
}
