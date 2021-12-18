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

	return calculateCommands(data, isPart1)
}

func calculateCommands(lines []string, isPart1 bool) int {
	var hpos int
	var depth int
	var aim int

	for _, line := range lines {
		cmd := parseCommand(line)

		if strings.Contains(line, "forward") {
			handleForward(&hpos, &depth, &aim, cmd, isPart1)
			// hpos += cmd
		} else if strings.Contains(line, "up") {
			handleUp(&depth, &aim, cmd, isPart1)
			// depth -= cmd
		} else if strings.Contains(line, "down") {
			handleDown(&depth, &aim, cmd, isPart1)
			// depth += cmd
		}
	}

	return hpos * depth
}

func handleDown(depth *int, aim *int, i int, isPart1 bool) {
	if isPart1 {
		*depth += i
		return
	}
	*aim += i
}

func handleUp(depth *int, aim *int, i int, isPart1 bool) {
	if isPart1 {
		*depth -= i
		return
	}
	*aim -= i
}

func handleForward(hpos *int, depth *int, aim *int, i int, isPart1 bool) {
	if isPart1 {
		*hpos += i
		return
	}

	*hpos += i
	*depth += *aim * i
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
