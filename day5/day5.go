package day5

import (
	"math"
)

func removeUnitPair(input string, unit rune) string {
	// TODO: optimize
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
	unit := rune(65)
	polymer := ""
	
	min := len(input)
	for i := 65; i <= 90; i++ {
		reduced := removeUnitPair(input, rune(i))
		result := React(reduced)
		
		if len(result) < min {
			min = len(result)
			unit = rune(i)
			polymer = result
		}
	}
	
	return unit, polymer
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
