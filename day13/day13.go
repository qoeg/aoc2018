package day13

import (
	"fmt"
	"sort"

	"github.com/qoeg/aoc2018/util"
)

type turn int

const (
	left turn = iota
	straight
	right
)

const (
	cartLeft = 60
	cartRight = 62
	cartUp = 94
	cartDown = 118
	curveLeftDownRightUp = 47
	curveLeftUpRightDown = 92
	intersection = 43
	lineEnd = 10
	space = 32
	straightHorz = 45
	straightVert = 124
)

// Grid represents the grid of tracks
type Grid struct {
	cells [][]rune
	Carts []*cart
	Crashes []util.Coordinate
}

// NewGrid parses the input and creates the grid
func NewGrid(input string) *Grid {
	var g = &Grid{}

	var x, y int
	for _, r := range input {
		if r == rune(10) {
			y++
			x = 0
			continue
		}

		if r == rune(cartLeft) || r == rune(cartRight) {
			c := &cart{Pos: util.Coordinate{X:x,Y:y}, symbol: int(r)}
			g.Carts = append(g.Carts, c)
			r = rune(straightHorz)
		}

		if r == rune(cartUp) || r == rune(cartDown) {
			c := &cart{Pos: util.Coordinate{X:x,Y:y}, symbol: int(r)}
			g.Carts = append(g.Carts, c)
			r = rune(straightVert)
		}

		if y == 0 {
			g.cells = append(g.cells, []rune{})
		}

		g.cells[x] = append(g.cells[x], r)
		x++
	}

	return g
}

func (g *Grid) cart(c *cart) *cart {
	for _, ct := range g.Carts {
		if c == ct {
			return ct
		}
	}
	return nil
}

func (g *Grid) cartAt(coord util.Coordinate, exclude *cart) *cart {
	for _, c := range g.Carts {
		if c == exclude {
			continue
		}

		if c.Pos == coord {
			return c
		}
	}
	return nil
}

func (g *Grid) crash(coord util.Coordinate) (util.Coordinate, bool) {
	for _, c := range g.Crashes {
		if c == coord {
			return c, true
		}
	}
	return util.Coordinate{}, false
}

func (g *Grid) remove(crashed *cart) {
	for i, c := range g.Carts {
		if c == crashed {
			g.Carts = append(g.Carts[:i], g.Carts[i+1:]...)
		}
	}
}

// Render draws the grid to the console
func (g *Grid) Render(crashes bool) {
	var x, y int
	for ; y < len(g.cells[0]); y++ {
		for x = 0; x < len(g.cells); x++ {
			sym := g.cells[x][y]
			if c := g.cartAt(util.Coordinate{X:x,Y:y}, nil); c != nil {
				sym = rune(c.symbol)
			}
			if _, ok := g.crash(util.Coordinate{X:x,Y:y}); ok && crashes {
				sym = 'X'
			}
			fmt.Print(string(sym))
		}
		fmt.Println()
	}
}

// Move moves all of the carts for 1 tick
func (g *Grid) Move() bool {
	sort.Slice(g.Carts, func(i, j int) bool {
		if g.Carts[i].Pos.Y == g.Carts[j].Pos.Y {
			return g.Carts[i].Pos.X < g.Carts[j].Pos.X
		}
		return g.Carts[i].Pos.Y < g.Carts[j].Pos.Y
	})

	var carts = make([]*cart, len(g.Carts))
	for i := range g.Carts {
		carts[i] = g.Carts[i]
	}

	for i := 0; i < len(carts); i++ {
		c := g.cart(carts[i])
		if c == nil {
			continue
		}

		c.move(g)

		if cc := g.cartAt(c.Pos, c); cc != nil {
			g.Crashes = append(g.Crashes, c.Pos)

			g.remove(c)
			g.remove(cc)
		}
	}

	return len(g.Carts) > 1
}
