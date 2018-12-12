package day11

import (
	"fmt"
)

// Coordinate is a value with x and y parts
type Coordinate struct {
	x int
	y int
}

// Square represents a square of Coordinates in a Grid
type Square struct {
	Coordinate
	Size int
	Sum int
}

// Grid representes an x,y plane of power levels
type Grid struct {
	SN int
	Cells map[int]map[int]int
	size int
}

// NewGrid creates a fully populated grid
func NewGrid(sn, size int) Grid {
	g := Grid{
		SN: sn,
		Cells: make(map[int]map[int]int),
		size: size,
	}
	
	for x := 1; x <= size; x++ {
		if _, ok := g.Cells[x]; !ok {
			g.Cells[x] = make(map[int]int, size)
		}

		for y := 1; y <= size; y++ {
			g.Cells[x][y] = powerLevel(x, y, sn)
		}
	}
	
	return g
}

func hundredth(n int) int {
	return int(n / 100) % 10
}

func powerLevel(x, y, sn int) int {
	rackID := x + 10
	pl := rackID * y
	pl += sn
	pl *= rackID
	return hundredth(pl) - 5
}

func sumSquare(g Grid, sq Square) int {
	sum := 0
	maxX, maxY := (sq.x + sq.Size), (sq.y + sq.Size)
	for x := sq.x; x < maxX; x++ {
		for y := sq.y; y < maxY; y++ {
			sum += g.Cells[x][y]
		}
	}
	return sum
}

// FindAnyLargestSquare finds the largest square of any size
func FindAnyLargestSquare(g Grid) Square {
	ch := make(chan Square, g.size)
	sem := make(chan struct{}, 1)

	for size := 1; size <= g.size; size++ {
		sem <- struct{}{}
		go func(size int) {
			sq := FindLargestSquare(g, size)
			fmt.Printf("Completed size: %d\n", size)
			<-sem
			ch<- sq
		}(size)
	}

	max := Square{}
	for i := 0; i < g.size; i++ {
		sq := <- ch
		if sq.Sum > max.Sum {
			max = sq
		}
	}

	return max
}

// FindLargestSquare finds the largest square fo a specified size
func FindLargestSquare(g Grid, size int) Square {
	ch := make(chan Square, 50)
	sem := make(chan struct{}, 50)

	bound := g.size-size
	for x := 1; x <= bound; x++ {
		for y := 1; y <= bound; y++ {
			sem <- struct{}{}
			go func(x, y int) {
				s := sumSquare(g, Square{Coordinate{x, y}, size, 0})
				<-sem
				ch<- Square{Coordinate{x, y}, size, s}
			}(x, y)
		}
	}

	max := Square{}
	total := (bound * bound)
	for i := 0; i < total; i++ {
		sq := <- ch
		if sq.Sum > max.Sum {
			max = sq
		}
	}

	return max
}
