package day5

import (
	"testing"
)

var input = "dabAcCaCBAcCcaDA"

func TestReact(t *testing.T) {
	output := React(input)

	if len(output) != 10 {
		t.Logf("Expected 10, but got %d", len(output))
		t.Fail()
	}
}

func TestFindOptimalUnit(t *testing.T) {
	unit, result := FindOptimalUnit(input)

	if unit != 67 {
		t.Logf("Expected 67, but got %d", unit)
		t.Fail()
	}

	if len(result) != 4 {
		t.Logf("Expected 4, but got %d", len(result))
		t.Fail()
	}
}
