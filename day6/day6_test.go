package day6

import (
	"fmt"
	"testing"
)

var input = []coordinate{
	{1,1},
	{1,6},
	{8,3},
	{3,4},
	{5,5},
	{8,9},
}

func TestMinMax(t *testing.T) {
	minX := minX(input)
	maxX := maxX(input)
	minY := minY(input)
	maxY := maxY(input)

	if minX != 1 && maxX != 8 {
		t.Logf("Expected 1 and 8, but got %d and %d", minX, maxX)
		t.Fail()
	}

	if minY != 1 && maxY != 9 {
		t.Logf("Expected 1 and 9, but got %d and %d", minY, maxY)
		t.Fail()
	}
}

func TestGetGrid(t *testing.T) {
	minX := minX(input)
	maxX := maxX(input)
	minY := minY(input)
	maxY := maxY(input)
	
	grid := getGrid(input)
	for y := minY; y <= maxY; y++ {
		fmt.Printf("%d: [", y)
		for x := minX; x <= maxX; x++ {
			fmt.Printf("%d:%v ", x, grid[x][y])
		}
		fmt.Println("]")
	}
}

func TestLargestAreaSize(t *testing.T) {
	size := LargestAreaSize(input)

	if size != 17 {
		t.Logf("Expected 17, but got %d", size)
		t.Fail()
	}
}

func TestSafeRegionSize(t *testing.T) {
	size := SafeRegionSize(input, 32)

	if size != 16 {
		t.Logf("Expected 16, but got %d", size)
		t.Fail()
	}
}
