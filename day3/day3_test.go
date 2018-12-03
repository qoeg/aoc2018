package day3

import (
	"testing"
)

func TestParseInput1(t *testing.T) {
	c := ParseInput("#1 @ 596,731: 11x27")

	if c.ID != 1 ||
	c.x != 596 ||
	c.y != 731 ||
	c.width != 11 ||
	c.height != 27 {
		t.Logf("Expected different properties for claim")
		t.Fail()
	}
}

func TestClaimsIntersect(t *testing.T) {
	c1 := ParseInput("#1267 @ 95,29: 17x28")
	c2 := ParseInput("#591 @ 97,20: 21x21")
	
	_, found := DoClaimsIntersect(c1, c2)

	if !found {
		t.Logf("Expected intersection but none found")
		t.Fail()
	}

	// assert bounds
}

func TestGetNumIntersectingSqIn(t *testing.T) {
	c1 := ParseInput("#1267 @ 95,29: 17x28")
	c2 := ParseInput("#591 @ 97,20: 21x21")

	count := GetNumIntersectingSqIn([]Claim{c1, c2})
	if count != 180 {
		t.Logf("Expected 180 intersecting, but got %d", count)
		t.Fail()
	}
}
