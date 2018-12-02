package main

import (
	"fmt"

	"github.com/qoeg/aoc2018/day1"
	"github.com/qoeg/aoc2018/day2"
)

func main() {
	d2()
}

func d2() {
	fmt.Print("Day 2\n")

	count2 := day2.CountContainsTwo(day2.Input)
	count3 := day2.CountContainsThree(day2.Input)
	fmt.Printf("Puzzle 1 Answer: %d\n", count2 * count3)

	correctID := day2.GetCorrectID(day2.Input)
	fmt.Printf("Puzzle 2 Answer: %s\n", correctID)
}

func d1() {
	fmt.Print("Day 1\n")

	d1p1 := day1.CalculateFreq(0, day1.Input)
	fmt.Printf("Puzzle 1 Answer: %d\n", d1p1)

	d1p2 := day1.CalibrateFreq(0, day1.Input)
	fmt.Printf("Puzzle 2 Answer: %d\n", d1p2)
}
