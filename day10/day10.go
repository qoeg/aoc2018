package day10

import (
	"fmt"
	"math"
	"time"
)

type coordinate struct {
	x int
	y int
}

// Point represents a single x,y point with velocity.
type Point struct {
	coordinate
	vX int
	vY int
}

// Grid represents a planar space to print points.
type Grid struct {
	minX int
	minY int
	maxX int
	maxY int
}

// NewGrid creates a grid for printing points.
func NewGrid(minX, minY, maxX, maxY int) Grid {
	return Grid{minX, minY, maxX, maxY}
}

func minX(points []*Point) int {
	min := math.MaxInt32
	for _, c := range points {
		if c.x < min {
			min = c.x
		}
	}
	return min
}

func maxX(points []*Point) int {
	max := 0
	for _, c := range points {
		if c.x > max {
			max = c.x
		}
	}
	return max
}

func minY(points []*Point) int {
	min := math.MaxInt32
	for _, c := range points {
		if c.y < min {
			min = c.y
		}
	}
	return min
}

func maxY(points []*Point) int {
	max := 0
	for _, c := range points {
		if c.y > max {
			max = c.y
		}
	}
	return max
}

func makeMap(points []*Point) map[coordinate]*Point {
	result := make(map[coordinate]*Point, len(points))
	for _, point := range points {
		result[point.coordinate] = point
	}
	return result
}

func parsePoints(input []string) []*Point {
	results := []*Point{}
	for i := range input {
		var x, y, vX, vY int
		fmt.Sscanf(input[i], "position=<%d, %d> velocity=<%d, %d>", &x, &y, &vX, &vY)

		results = append(results, &Point{coordinate{x, y}, vX, vY})
	}
	return results
}

func printPoints(g Grid, points []*Point) {
	pm := makeMap(points)
	for y := g.minY; y <= g.maxY; y++ {
		for x := g.minX; x <= g.maxX; x++ {
			if pm[coordinate{x, y}] != nil {
				fmt.Print("#")
				continue
			}

			if x == 0 {
				fmt.Print("|")
				continue
			}

			if y == 0 {
				fmt.Print("-")
				continue
			}

			fmt.Print(".")
		}
		fmt.Println()
	}
}

func updatePoints(points []*Point) []*Point {
	results := []*Point{}
	for _, p := range points {
		x := (p.x + p.vX)
		y := (p.y + p.vY)
		results = append(results, &Point{
			coordinate{x, y},
			p.vX,
			p.vY,
		})
	}
	return results
}

// Run runs the progression of the points parsed from the input.
// It prints to the console the points on an automatically bounded grid.
func Run(input []string, skip, end int) {
	points := parsePoints(input)

	minX := minX(points)
	maxX := maxX(points)
	minY := minY(points)
	maxY := maxY(points)
	g := Grid{minX, minY, maxX, maxY}

	RunOnGrid(g, input, skip, end)
}

// RunOnGrid runs the progression of the points parsed from the input.
// It prints to the console the points on a predefined grid. Specify a
// skip and end intiger to skip time and stop at a particular second.
func RunOnGrid(g Grid, input []string, skip, end int) {
	p := parsePoints(input)

	if skip == 0 {
		fmt.Println("Initially:")
		printPoints(g, p)
		fmt.Println()

		time.Sleep(time.Millisecond * 1000)
	}

	count := 0
	next := func() bool {
		if end < 0 {
			return false
		}
		return count <= end
	}
	
	for next() {
		if count > skip {
			fmt.Printf("After seconds: %d\n", count)
			printPoints(g, p)
			fmt.Println()

			time.Sleep(time.Millisecond * 1000)
		}

		p = updatePoints(p)
		count++
	}
}
