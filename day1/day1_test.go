package day1

import (
	"fmt"
	"testing"
)

var exampleAA = []int{1, -2, 3, 1}

var exampleBA = []int{1, -1}
var exampleBB = []int{3, 3, 4, -2, -4}
var exampleBC = []int{-6, 3, 8, 5, -6}
var exampleBD = []int{7, 7, -2, -7, -4}

func TestCalculateFreq_ExampleAA(t *testing.T) {
	result := CalculateFreq(0, exampleAA)
	if result != 3 {
		t.Logf("Expected frequency %d, but got %d\n", 3, result)
		t.Fail()
	}
}

func TestCalibrateFreq_ExampleBB(t *testing.T) {
	result := CalibrateFreq(0, exampleBB)
	fmt.Printf("*** Result: %d\n", result)

	if result != 10 {
		t.Logf("Expected frequency %d, but got %d\n", 10, result)
		t.Fail()
	}
}

func TestCalibrateFreq_ExampleBC(t *testing.T) {
	result := CalibrateFreq(0, exampleBC)
	fmt.Printf("*** Result: %d\n", result)

	if result != 5 {
		t.Logf("Expected frequency %d, but got %d\n", 5, result)
		t.Fail()
	}
}

func TestCalibrateFreq_ExampleBD(t *testing.T) {
	result := CalibrateFreq(0, exampleBD)
	fmt.Printf("*** Result: %d\n", result)

	if result != 14 {
		t.Logf("Expected frequency %d, but got %d\n", 14, result)
		t.Fail()
	}
}
