package day12

import (
	"fmt"
)

// Formula represents a formula for the next generation
type Formula struct {
	Pattern string
	Product string
}

func parseFormula(input string) Formula {
	var m, p string
	fmt.Sscanf(input, "%s => %s", &m, &p)

	return Formula{m, p}
}

func doGeneration(state string, formulas []Formula) string {
	result := state

	for i := 0; i < len(state)-5; i++ {
		next := '.'
		for _, f := range formulas {
			if state[i:i+5] == f.Pattern {
				next = rune(f.Product[0])
				break
			}
		}

		r := []rune(result)
		r[i+2] = next
		result = string(r)
	}

	return result
}

// Formulas gets the structured formulas from the input strings
func Formulas(inputs []string) []Formula {
	formulas := make([]Formula, len(inputs))
	for i := 0; i < len(formulas); i++ {
		formulas[i] = parseFormula(inputs[i])
	}
	return formulas
}

// DoGenerations runs through a number of plant generations from an initial state
// and collections of formulas. Use the print argument to print the lines to console.
func DoGenerations(initial string, forms []Formula, num int, print bool) (int, int) {
	s := "....." + initial + "....."

	for i := 0; i < num; i++ {
		s += "."
		s = doGeneration(s, forms)
		if print {
			fmt.Printf("%d: %s\n", i+1, s)
		}
	}

	var sum, count int
	for i := 0; i < len(s); i++ {
		if rune(s[i]) == rune('#') {
			sum += i-5
			count++
		}
	}

	return sum, count
}
