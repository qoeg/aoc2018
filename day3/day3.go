package day3

import (
	"fmt"
	"math"
)

// Extent ...
type Extent struct {
	x int
	y int
	width int
	height int
}

// Claim ...
type Claim struct {
	Code string
	ID int
	Extent
}

// ParseInput ...
func ParseInput(code string) Claim {
	var id, x, y, width, height int
	
	fmt.Sscanf(code, "#%d @ %d,%d: %dx%d", &id, &x, &y, &width, &height)

	return Claim{
		code,
		id,
		Extent{
			x,
			y,
			width,
			height,
		},
	}
}

// DoClaimsIntersect ...
func DoClaimsIntersect(c1 Claim, c2 Claim) (Extent, bool) {
	if c1.x + c1.width <= c2.x || c2.x + c2.width <= c1.x {
		return Extent{}, false
	}

	if c1.y + c1.height <= c2.y || c2.y + c2.height <= c1.y {
		return Extent{}, false
	}

	leftBound := math.Max(float64(c1.x), float64(c2.x))
	rightBound := math.Min(float64(c1.x + c1.width), float64(c2.x + c2.width))

	topBound := math.Max(float64(c1.y), float64(c2.y))
	bottomBound := math.Min(float64(c1.y + c1.height), float64(c2.y + c2.height))

	intersection := Extent{
		int(leftBound),
		int(topBound),
		int(rightBound - leftBound),
		int(bottomBound - topBound),
	}

	return intersection, true
}

// GetNumIntersectingSqIn ...
func GetNumIntersectingSqIn(claims []Claim) int {
	cells := map[int]map[int]int{}

	count := 0
	for _, claim := range claims {
		for x := claim.x; x < claim.x+claim.width; x++ {
			for y := claim.y; y < claim.y+claim.height; y++ {
				if cells[x] == nil {
					cells[x] = map[int]int{}
				}

				cells[x][y]++

				if cells[x][y] == 2 {
					count++
				}
			}
		} 
	}

	return count
}

// GetCleanClaim ...
func GetCleanClaim(claims []Claim) Claim {
	for _, claim1 := range claims {
		clean := true
		for _, claim2 := range claims {
			if claim1.ID == claim2.ID {
				continue
			}

			if _, yes := DoClaimsIntersect(claim1, claim2); yes {
				clean = false
			}
		}

		if clean {
			return claim1
		}
	}

	return Claim{}
}
