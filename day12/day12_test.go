package day12

import (
	"testing"
)

var state = "#..#.#..##......###...###"

var input = []string{
	"...## => #",
	"..#.. => #",
	".#... => #",
	".#.#. => #",
	".#.## => #",
	".##.. => #",
	".#### => #",
	"#.#.# => #",
	"#.### => #",
	"##.#. => #",
	"##.## => #",
	"###.. => #",
	"###.# => #",
	"####. => #",
}

func Test_parseFormula(t *testing.T) {
	f := parseFormula(input[0])

	if f.Pattern != "...##" {
		t.Errorf("Expected f.Pattern = %s, got %s", "...##", f.Pattern)
	}

	if f.Product != "#" {
		t.Errorf("Expected f.product = %s, got %s", "#", f.Product)
	}
}

func Test_DoGenerations(t *testing.T) {
	sum, _ := DoGenerations(state, Formulas(input), 20, false)

	if sum != 325 {
		t.Errorf("Expected sum = %d, got %d", 325, sum)
	}
}
