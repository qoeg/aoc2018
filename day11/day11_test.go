package day11

import (
	"testing"
)

var input = []string{}

func testPowerLevel(x, y, sn, exp int, t *testing.T) {
	pl := powerLevel(x, y, sn)
	
	if pl != exp {
		t.Errorf("Expected pl = %d, got %d", exp, pl)
	}
}
func Test_powerLevel4(t *testing.T) {
	testPowerLevel(3, 5, 8, 4, t)
}

func Test_powerLevel57(t *testing.T) {
	testPowerLevel(122, 79, 57, -5, t)
}

func Test_powerLevel39(t *testing.T) {
	testPowerLevel(217, 196, 39, 0, t)
}

func Test_powerLevel71(t *testing.T) {
	testPowerLevel(101, 153, 71, 4, t)
}

func Test_FindLargestSquare18(t *testing.T) {
	g := NewGrid(18, 300)
	s, p := FindLargestSquare(g, 3)

	if (s.Coordinate != Coordinate{33, 45}) {
		t.Errorf("Expected pl = %d, got %d", Coordinate{33, 45}, s.Coordinate)
	}

	if (p != 29) {
		t.Errorf("Expected pl = %d, got %d", 29, p)
	}
}

func Test_FindLargestSquare42(t *testing.T) {
	g := NewGrid(42, 300)
	s, p := FindLargestSquare(g, 3)

	if (s.Coordinate != Coordinate{21, 61}) {
		t.Errorf("Expected pl = %d, got %d", Coordinate{21, 61}, s.Coordinate)
	}

	if (p != 30) {
		t.Errorf("Expected pl = %d, got %d", 30, p)
	} 
}
