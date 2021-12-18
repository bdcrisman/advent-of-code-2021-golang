package utils

import (
	"bufio"
	"io/ioutil"
	"os"
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

// Reads file in as a string
func ReadAsString(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Reads file in as string array.
// https://stackoverflow.com/questions/5884154/read-text-file-into-string-array-and-write
func ReadAsStringArray(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
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
