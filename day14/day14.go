package day14

import (
	"fmt"
	"strconv"
	"strings"
)

type game struct {
	scores []byte
	elf1 int
	elf2 int
}

func newGame() *game {
	return &game{[]byte{'3','7'}, 0, 1}
}

func (g *game) print(newline bool) {
	for i, s := range g.scores {
		switch i {
		case g.elf1:
			fmt.Printf("(%d)", s-'0')
		case g.elf2:
			fmt.Printf("[%d]", s-'0')
		default:
			fmt.Printf(" %d ", s-'0')
		}
	}
	if newline {
		fmt.Println()
	}
}

func (g *game) makeMore() {
	score := []byte(strconv.Itoa(int(g.scores[g.elf1]-'0' + g.scores[g.elf2]-'0')))
	g.scores = append(g.scores, score...)
}

func (g *game) pickNew() {
	g.elf1 = (g.elf1 + 1 + int(g.scores[g.elf1]-'0')) % len(g.scores)
	g.elf2 = (g.elf2 + 1 + int(g.scores[g.elf2]-'0')) % len(g.scores)
}

func (g *game) run(stopAt func(*game) bool, print bool) {
	if print {
		g.print(true)
	}

	for !stopAt(g) {
		g.makeMore()
		g.pickNew()

		if print {
			g.print(true)
		}
	}
}

// Answer1 gets the answer for the first part
// skip is the number of scores to skip, and take is the number to return
func Answer1(skip, take int) string {
	g := newGame()
	stopAt := func(g *game) bool {
		return len(g.scores) >= skip+take
	}

	g.run(stopAt, false)

	return string(g.scores[skip:skip+take])
}

// Answer2 gets the answer for the second part
// match is the string of scores to match, and after is the number of scores to generate
func Answer2(match string, after int) int {
	g := newGame()
	stopAt := func(g *game) bool {
		return len(g.scores) >= after
	}

	g.run(stopAt, false)

	return strings.Index(string(g.scores), match)
}
