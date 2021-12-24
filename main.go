package main

import (
	"fmt"
	"path/filepath"

	"github.com/bdcrisman/advent-of-code/2021/day1"
	"github.com/bdcrisman/advent-of-code/2021/day2"
	"github.com/bdcrisman/advent-of-code/2021/day3"
)

func main() {
	////////////////////////////////////////////////////////////////////////////////////////////////
	d1p1 := day1.Day1(filepath.Join("day1", "input"), true)
	fmt.Printf("Day 1 Part 1: %d\n", d1p1)
	d1p2 := day1.Day1(filepath.Join("day1", "input"), false)
	fmt.Printf("Day 1 Part 2: %d\n", d1p2)

	////////////////////////////////////////////////////////////////////////////////////////////////
	d2p1 := day2.Day2(filepath.Join("day2", "input"), true)
	fmt.Printf("Day 2 Part 1: %d\n", d2p1)
	d2p2 := day2.Day2(filepath.Join("day2", "input"), false)
	fmt.Printf("Day 2 Part 1: %d\n", d2p2)

	////////////////////////////////////////////////////////////////////////////////////////////////
	d3p1 := day3.Day3(filepath.Join("day3", "input"), true)
	fmt.Printf("Day 3 Part 1: %d\n", d3p1)
	d3p2 := day3.Day3(filepath.Join("day3", "input"), false)
	fmt.Printf("Day 3 Part 2: %d\n", d3p2)
}
