package main

import (
	"fmt"

	"github.com/bdcrisman/advent-of-code/2021/day1"
	"github.com/bdcrisman/advent-of-code/2021/day2"
)

func main() {
	d1p1 := day1.Day1("input", true)
	fmt.Printf("Day 1 Part 1 :: Number increases: %d\n", d1p1)
	d1p2 := day1.Day1("input", false)
	fmt.Printf("Day 1 Part 2 :: Number increases: %d\n", d1p2)

	d2p1 := day2.Day2("input", true)
	fmt.Printf("Day 2 Part 1: %d\n", d2p1)
	d2p2 := day2.Day2("input", false)
	fmt.Printf("Day 2 Part 1: %d\n", d2p2)
}
