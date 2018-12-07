package main

import (
	"fmt"

	"github.com/qoeg/aoc2018/day1"
	"github.com/qoeg/aoc2018/day2"
	"github.com/qoeg/aoc2018/day3"
	"github.com/qoeg/aoc2018/day4"
	"github.com/qoeg/aoc2018/day5"
	"github.com/qoeg/aoc2018/day6"
)

func main() {
	d6()
}

func d6() {
	fmt.Print("Day 6\n")

	areaSize := day6.LargestAreaSize(day6.Input)
	fmt.Printf("Puzzle 1 Answer: %v\n", areaSize)

	regSize := day6.SafeRegionSize(day6.Input, 10000)
	fmt.Printf("Puzzle 2 Answer: %v\n", regSize)
}

func d5() {
	fmt.Print("Day 5\n")

	output := day5.React(day5.Input)
	fmt.Printf("Puzzle 1 Answer: %v\n", len(output))

	unit, shortest := day5.FindOptimalUnit(day5.Input)
	fmt.Printf("Puzzle 2 Answer: %v (with unit %s/%s)\n", len(shortest), string(unit), string(unit+32))
}

func d4() {
	fmt.Print("Day 4\n")

	guards := day4.GetGuardSchedules(day4.Input)

	bestGuard := day4.MostSleepy(guards)
	minute, _ := day4.MostCommonMinute(bestGuard)
	fmt.Printf("Puzzle 1 Answer: %v\n", bestGuard.ID * minute)

	var id, bestMinute, maxFreq int
	for _, g := range guards {
		m, f := day4.MostCommonMinute(g)
		if f > maxFreq {
			maxFreq = f
			bestMinute = m
			id = g.ID
		}
	}
	fmt.Printf("Puzzle 2 Answer: %v\n", id * bestMinute)
}

func d3() {
	fmt.Print("Day 3\n")

	claims := []day3.Claim{}
	for _, input := range day3.Input {
		claims = append(claims, day3.ParseInput(input))
	}

	answer1 := day3.GetNumIntersectingSqIn(claims)
	fmt.Printf("Puzzle 1 Answer: %v\n", answer1)

	answer2 := day3.GetCleanClaim(claims)
	fmt.Printf("Puzzle 2 Answer: %v\n", answer2.ID)
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
