package day5

import (
	"math"
)

func removeUnitPair(input string, unit rune) string {
	runes := []rune(input)
	for i := 0; i < len(runes); i++ {
		if runes[i] == unit || runes[i] == unit+32 {
			runes = append(runes[:i], runes[i+1:]...)
			i--
		}
	}

	return string(runes)
}

// FindOptimalUnit tests all of the possible units by reacting them then, finding the best one
func FindOptimalUnit(input string) (rune, string) {
	type result struct {
		unit rune
		polymer string
	}
	results := make(chan result)

	for i := 65; i <= 90; i++ {
		go func(unit int){
			primed := removeUnitPair(input, rune(unit))
			reduced := React(primed)

			results <- result{
				unit: rune(unit),
				polymer: reduced,
			}
		}(i)
	}

	cnt := 0
	opt := result{}
	min := len(input)

	for r := range results {
		if len(r.polymer) < min {
			min = len(r.polymer)
			opt = r
		}

		cnt++
		if cnt >= 26 {
			break
		}
	}
	
	return opt.unit, opt.polymer
}

// React causes the unit reaction which reduces the polymer
func React(input string) string {
	runes := []rune(input)
	for i := 1; i < len(runes); i++ {
		if math.Abs(float64(runes[i]) - float64(runes[i-1])) == 32 {
			runes = append(runes[:i-1], runes[i+1:]...)
			i = 0
		}
	}

	return string(runes)
}
