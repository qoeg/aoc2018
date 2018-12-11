package day11

import (
	"fmt"
)

type Coordinate struct {
	x int
	y int
}

type Square struct {
	Coordinate
	Size int
	sum int
}

type Grid struct {
	SN int
	Cells map[int]map[int]int
	size int
}

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

func FindAnyLargestSquare(g Grid) Square {
	ch := make(chan Square, g.size)

	for size := 1; size <= g.size; size++ {
		go func(size int) {
			fmt.Printf("Processing size: %d\n", size)
			sq, _ := FindLargestSquare(g, size)
			fmt.Printf("Completed size: %d\n", size)
			ch<- sq
		}(size)
	}

	max := Square{}
	for i := 0; i < g.size; i++ {
		sq := <- ch
		if sq.sum > max.sum {
			max = sq
		}
	}

	return max
}

func FindLargestSquare(g Grid, size int) (Square, int) {
	sum := 0
	max := Square{Coordinate{1, 1}, size, 0}
	for x := 1; x <= 300-size; x++ {
		for y := 1; y <= 300-size; y++ {
			s := sumSquare(g, Square{Coordinate{x, y}, size, 0})
			if s > sum {
				max = Square{Coordinate{x, y}, size, s}
				sum = s
			}
		}
	}
	return max, sum
}
