package main

import (
	"fmt"

	"github.com/qoeg/aoc2018/day1"
	"github.com/qoeg/aoc2018/day2"
	"github.com/qoeg/aoc2018/day3"
	"github.com/qoeg/aoc2018/day4"
	"github.com/qoeg/aoc2018/day5"
	"github.com/qoeg/aoc2018/day6"
	"github.com/qoeg/aoc2018/day7"
	"github.com/qoeg/aoc2018/day8"
	"github.com/qoeg/aoc2018/day9"
	"github.com/qoeg/aoc2018/day10"
	"github.com/qoeg/aoc2018/day11"
	"github.com/qoeg/aoc2018/day12"
	"github.com/qoeg/aoc2018/day13"
)

func main() {
	d13()
}

func d13() {
	fmt.Print("Day 13\n")

	g := day13.NewGrid(day13.Input)
	for g.Move() {}

	fmt.Printf("Puzzle 1 Answer: %v\n", g.Crashes[0])
	fmt.Printf("Puzzle 2 Answer: %v\n", g.Carts[0].Pos)
}

func d12() {
	fmt.Print("Day 12\n")

	frms := day12.Formulas(day12.Input)
	sum1, _ := day12.DoGenerations(day12.InitialState, frms, 20, true)
	fmt.Printf("Puzzle 1 Answer: %v\n", sum1)

	sum2, count2 := day12.DoGenerations(day12.InitialState, frms, 111, true)
	fmt.Printf("Puzzle 2 Answer: %v\n", ((50000000000-111)*count2)+sum2)
}

func d11() {
	fmt.Print("Day 11\n")

	g := day11.NewGrid(day11.Input, 300)
	sq1 := day11.FindLargestSquare(g, 3)
	fmt.Printf("Puzzle 1 Answer: %v\n", sq1.Coordinate)

	sq2 := day11.FindAnyLargestSquare(g)
	fmt.Printf("Puzzle 2 Answer: %v\n", sq2)
}

func d10() {
	fmt.Print("Day 10\n")

	fmt.Println("Puzzle Answer:")
	day10.RunOnGrid(day10.NewGrid(134, 116, 206, 134), day10.Input, 10123, 10124)
}

func d9() {
	fmt.Print("Day 9\n")

	game1 := day9.NewGameStr(day9.Input)
	fmt.Printf("Puzzle 1 Answer: %v\n", day9.GetHighScoreFast(game1))

	game2 := day9.NewGameStr(day9.Input)
	game2.LastMarble *= 100
	fmt.Printf("Puzzle 2 Answer: %v\n", day9.GetHighScoreFast(game2))
}

func d8() {
	fmt.Print("Day 8\n")

	node, _ := day8.ParseNode(day8.Input)
	sum := day8.SumMetadata(node)
	fmt.Printf("Puzzle 1 Answer: %v\n", sum)

	valsum := day8.SumValues(node)
	fmt.Printf("Puzzle 2 Answer: %v\n", valsum)
}

func d7() {
	fmt.Print("Day 7\n")

	steps, _ := day7.OrderSteps(day7.Input, 1, 0)
	fmt.Printf("Puzzle 1 Answer: %v\n", steps)

	_, seconds := day7.OrderSteps(day7.Input, 5, 60)
	fmt.Printf("Puzzle 2 Answer: %v\n", seconds)
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
