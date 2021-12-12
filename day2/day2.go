package day2

import (
	"path/filepath"
	"strconv"
	"strings"

	"github.com/bdcrisman/advent-of-code/2021/utils"
)

func Day2(inputPath string, isPart1 bool) int {
	path := filepath.Join("day2", inputPath)
	data, err := utils.ReadAsStringArray(path)
	if err != nil {
		println(err.Error())
		return 0
	}

	return calculateCommands(data)
}

func calculateCommands(lines []string) int {
	var hpos int
	var depth int

	for _, line := range lines {
		cmd := parseCommand(line)

		if strings.Contains(line, "forward") {
			hpos += cmd
		} else if strings.Contains(line, "up") {
			depth -= cmd
		} else if strings.Contains(line, "down") {
			depth += cmd
		}
	}

	return hpos * depth
}

func parseCommand(line string) int {
	s := strings.Split(line, " ")
	if len(s) < 2 {
		return 0
	}

	v, err := strconv.Atoi(s[1])
	if err != nil {
		return 0
	}
	return v
}
