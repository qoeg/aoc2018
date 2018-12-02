package main

import (
	"fmt"

	"github.com/qoeg/aoc2018/day1"
)

func main() {
	d1()
}

func d1() {
	d1p1 := day1.CalculateFreq(0, day1.Input)
	fmt.Printf("Day 1, Puzzle 1 Answer: %d\n", d1p1)

	d1p2 := day1.CalibrateFreq(0, day1.Input)
	fmt.Printf("Day 1, Puzzle 2 Answer: %d\n", d1p2)
}
