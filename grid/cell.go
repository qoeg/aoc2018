package grid

import (

)

type Cell struct {
	index int
	mark rune
	pos Coordinate
}

func NewCell(mark rune, pos Coordinate) Cell {
	return Cell{0, mark, pos}
}

type Renderable interface {
	Pos() Coordinate
	String() string
}

func (c Cell) Mark() rune {
	return c.mark
}

func (c Cell) Pos() Coordinate {
	return c.pos
}

func (c Cell) String() string {
	return string(c.mark)
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
