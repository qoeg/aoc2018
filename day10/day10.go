package day10

import (
	"time"
	"fmt"
	"math"
)

type grid struct {
	minX int
	minY int
	maxX int
	maxY int
}

type coordinate struct {
	x int
	y int
}

type Point struct {
	coordinate
	vX int
	vY int
}

func minX(points map[coordinate]*Point) int {
	min := math.MaxInt32
	for _, c := range points {
		if c == nil { continue }
		if c.x < min {
			min = c.x
		}
	}
	return min
}

func maxX(points map[coordinate]*Point) int {
	max := 0
	for _, c := range points {
		if c == nil { continue }
		if c.x > max {
			max = c.x
		}
	}
	return max
}

func minY(points map[coordinate]*Point) int {
	min := math.MaxInt32
	for _, c := range points {
		if c == nil { continue }
		if c.y < min {
			min = c.y
		}
	}
	return min
}

func maxY(points map[coordinate]*Point) int {
	max := 0
	for _, c := range points {
		if c == nil { continue }
		if c.y > max {
			max = c.y
		}
	}
	return max
}

func parsePoints(input []string) map[coordinate]*Point {
	results := make(map[coordinate]*Point, len(input))
	for i := range input {
		var x, y, vX, vY int
		fmt.Sscanf(input[i], "position=<%d, %d> velocity=<%d, %d>", &x, &y, &vX, &vY)

		results[coordinate{x, y}] = &Point{coordinate{x, y}, vX, vY}
	}
	return results
}

func printPoints(g grid, points map[coordinate]*Point) {
	for y := g.minY; y <= g.maxY; y++ {
		for x := g.minX; x <= g.maxX; x++ {
			if points[coordinate{x, y}] != nil {
				fmt.Print("#")
				fmt.Printf("(%d,%d[%d,%d])", x, y, points[coordinate{x,y}].vX, points[coordinate{x,y}].vY)
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

func updatePoints(points map[coordinate]*Point) {
	for coord := range points {
		if points[coord] == nil {continue}

		x := (points[coord].x + points[coord].vX)
		y := (points[coord].y + points[coord].vY)
		p := &Point{
			coordinate{x, y},
			points[coord].vX,
			points[coord].vY,
		}
		points[coordinate{p.x, p.y}] = p
		points[coord] = nil
	}
}

func Run(input []string) {
	points := parsePoints(input)
	minX := minX(points)
	maxX := maxX(points)
	minY := minY(points)
	maxY := maxY(points)
	g := grid{minX, minY, maxX, maxY}
	count := 0
	
	fmt.Println("Initially:")
	for {
		printPoints(g, points)
		updatePoints(points)
		
		time.Sleep(time.Second * 2)
		count++
		fmt.Println()
		fmt.Printf("After seconds: %d\n", count)
	}
}
