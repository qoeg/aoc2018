package day13

import (
	"github.com/qoeg/aoc2018/util"
)

type cart struct {
	Pos util.Coordinate
	symbol int
	nextTurn turn
}

func (c *cart) move(g *Grid) {
	switch c.symbol {
	case cartLeft:
		switch next := g.cells[c.Pos.X-1][c.Pos.Y]; next {
		case rune(curveLeftDownRightUp):
			c.symbol = cartDown
		case rune(curveLeftUpRightDown):
			c.symbol = cartUp
		case rune(intersection):
			c.turn()
		}
		c.Pos = util.Coordinate{X:c.Pos.X-1,Y:c.Pos.Y}
	case cartDown:
		switch next := g.cells[c.Pos.X][c.Pos.Y+1]; next {
		case rune(curveLeftDownRightUp):
			c.symbol = cartLeft
		case rune(curveLeftUpRightDown):
			c.symbol = cartRight
		case rune(intersection):
			c.turn()
		}
		c.Pos = util.Coordinate{X:c.Pos.X,Y:c.Pos.Y+1}
	case cartRight:
		switch next := g.cells[c.Pos.X+1][c.Pos.Y]; next {
		case rune(curveLeftDownRightUp):
			c.symbol = cartUp
		case rune(curveLeftUpRightDown):
			c.symbol = cartDown
		case rune(intersection):
			c.turn()
		}
		c.Pos = util.Coordinate{X:c.Pos.X+1,Y:c.Pos.Y}
	case cartUp:
		switch next := g.cells[c.Pos.X][c.Pos.Y-1]; next {
		case rune(curveLeftDownRightUp):
			c.symbol = cartRight
		case rune(curveLeftUpRightDown):
			c.symbol = cartLeft
		case rune(intersection):
			c.turn()
		}
		c.Pos = util.Coordinate{X:c.Pos.X,Y:c.Pos.Y-1}
	}
}

func (c *cart) turn() {
	switch c.nextTurn {
	case left:
		switch c.symbol {
		case cartLeft:
			c.symbol = cartDown
		case cartDown:
			c.symbol = cartRight
		case cartRight:
			c.symbol = cartUp
		case cartUp:
			c.symbol = cartLeft
		}
	case right:
		switch c.symbol {
		case cartLeft:
			c.symbol = cartUp
		case cartUp:
			c.symbol = cartRight
		case cartRight:
			c.symbol = cartDown
		case cartDown:
			c.symbol = cartLeft
		}
	}
	c.nextTurn = (c.nextTurn + 1) % 3
}
