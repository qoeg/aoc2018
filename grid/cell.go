package grid

import (

)

type Cell struct {
	index int
	mark rune
	pos Coordinate
}

type Renderable interface {
	Mark() rune
	Pos() Coordinate
}

func (c Cell) Mark() rune {
	return c.mark
}

func (c Cell) Pos() Coordinate {
	return c.pos
}

type cellByReadOrder []Cell

func (c cellByReadOrder) Len() int {
	return len(c)
}

func (c cellByReadOrder) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c cellByReadOrder) Less(i, j int) bool {
	if c[i].pos.Y == c[j].pos.Y {
		return c[i].pos.X < c[j].pos.X
	}
	return c[i].pos.Y < c[j].pos.Y
}
