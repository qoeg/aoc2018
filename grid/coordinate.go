package grid

import (
	"math"
)

type Coordinate struct {
	X int
	Y int
}

type TwoDimensional interface {
	Pos() Coordinate
}

func (c Coordinate) Pos() Coordinate {
	return c
}

func MinX(coords []Coordinate) int {
	min := math.MaxInt32
	for _, c := range coords {
		if c.X < min {
			min = c.X
		}
	}
	return min
}

func MaxX(coords []Coordinate) int {
	max := 0
	for _, c := range coords {
		if c.X > max {
			max = c.X
		}
	}
	return max
}

func MinY(coords []Coordinate) int {
	min := math.MaxInt32
	for _, c := range coords {
		if c.Y < min {
			min = c.Y
		}
	}
	return min
}

func MaxY(coords []Coordinate) int {
	max := 0
	for _, c := range coords {
		if c.Y > max {
			max = c.Y
		}
	}
	return max
}
