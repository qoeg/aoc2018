package day13

import (
	"fmt"

	"github.com/qoeg/aoc2018/util"
)

const (
	cartRight = 62
	cartDown = 118
	curve1 = 47
	curve2 = 92
	intersection = 43
	lineEnd = 10
	space = 32
	straightHorz = 45
	straightVert = 124
)

type Grid struct {
	cells [][]rune
	carts []util.Coordinate
}

func NewGrid() {
	var g = Grid{}
	println(len(g.cells))
}

func parse(input string) Grid {
	var g = Grid{}

	var x, y int
	for _, r := range input {
		if len(g.cells) < x {
			g.cells = append(g.cells, []rune{})
		}

		g.cells[x][y] = r
		
		if  r == rune(cartRight) ||
			r == rune(cartDown) {
			g.carts = append(g.carts, util.Coordinate{X:x,Y:y})
		}

		if r == rune(10) {
			y++
		}

		x++
	}

	return g
}

func render(g Grid) {
	var x, y int
	for ; y < len(g.cells[0]); y++ {
		for x = 0; x < len(g.cells); x++ {
			fmt.Print(g.cells[x][y])
		}
		fmt.Println()
	}
}
