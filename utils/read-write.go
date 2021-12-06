package utils

import (
	"bufio"
	"io/ioutil"
	"strconv"
	"strings"
)

func TestEcho(s string) string {
	return s
}

func ReadAsIntArray(path string) ([]int, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	r := strings.NewReader(string(data))
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
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
