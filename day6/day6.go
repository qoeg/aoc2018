package day6

import (
	"math"
	"sort"
)

type coordinate struct {
	x int
	y int
}

func minX(coords []coordinate) int {
	min := math.MaxInt32
	for _, c := range coords {
		if c.x < min {
			min = c.x
		}
	}
	return min
}

func maxX(coords []coordinate) int {
	max := 0
	for _, c := range coords {
		if c.x > max {
			max = c.x
		}
	}
	return max
}

func minY(coords []coordinate) int {
	min := math.MaxInt32
	for _, c := range coords {
		if c.y < min {
			min = c.y
		}
	}
	return min
}

func maxY(coords []coordinate) int {
	max := 0
	for _, c := range coords {
		if c.y > max {
			max = c.y
		}
	}
	return max
}

func distance(a coordinate, b coordinate) int {
	return int(math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y)))
}

func getGrid(coords []coordinate) map[int]map[int]coordinate {
	minX := minX(coords)
	maxX := maxX(coords)
	minY := minY(coords)
	maxY := maxY(coords)

	grid := map[int]map[int]coordinate{}

	for x := minX; x <= maxX; x++ {
		grid[x] = map[int]coordinate{}

		for y := minY; y <= maxY; y++ {
			here := coordinate{x, y}
			min := distance(coordinate{minX, minY}, coordinate{maxX, maxY})
			counts := map[int]int{}

			for i := range coords {
				if coords[i] == here {
					grid[x][y] = coords[i]
					break
				}
				
				dist := distance(here, coords[i])
				counts[dist]++

				if dist <= min {
					grid[x][y] = coords[i]
					min = dist
				}
			}

			if len(counts) > 0 {
				var keys []int
				for k := range counts {
					keys = append(keys, k)
				}
				sort.Ints(keys)
				
				if counts[keys[0]] > 1 {
					grid[x][y] = coordinate{}
				}
			}
		}
	}

	return grid
}

// LargestAreaSize finds the largest area of cells closest to a coordinate
func LargestAreaSize(coords []coordinate) int {
	minX := minX(coords)
	maxX := maxX(coords)
	minY := minY(coords)
	maxY := maxY(coords)
	grid := getGrid(coords)

	counts := map[coordinate]int{}

	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			if (grid[x][y] == coordinate{}) {
				continue
			}
			
			if x == minX || x == maxX || y == minY || y == maxY {
				// Anti-border bias good enough for non-infinite areas
				continue
			}

			counts[grid[x][y]]++
		}
	}

	max := 0
	for _, count := range counts {
		if count > max {
			max = count
		}
	}

	return max
}

// SafeRegionSize finds the region size within a certain distance of all coordinates
func SafeRegionSize(coords []coordinate, dist int) int {
	minX := minX(coords)
	maxX := maxX(coords)
	minY := minY(coords)
	maxY := maxY(coords)

	var size int
	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			here := coordinate{x, y}

			sum := 0
			for i := range coords {
				sum += distance(here, coords[i])
			}

			if sum < dist {
				size++
			}
		}
	}

	return size
}
