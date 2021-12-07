package main

import (
	"fmt"

	"github.com/bdcrisman/advent-of-code/2021/day1"
)

func main() {
	d1p1 := day1.Day1("input", true)
	fmt.Printf("Day 1 Part 1 :: Number increases: %d\n", d1p1)

	d1p2 := day1.Day1("input", false)
	fmt.Printf("Day 1 Part 2 :: Number increases: %d\n", d1p2)
}
