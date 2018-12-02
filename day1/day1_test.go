package day1

import (
	"fmt"
	"testing"
)

func TestCalculateFreq_Example1(t *testing.T) {
	changes := []int{1, -2, 3, 1}

	result := CalculateFreq(0, changes)
	if result != 3 {
		t.Logf("Expected frequency %d, but got %d\n", 3, result)
		t.Fail()
	}
}

func TestCalibrateFreq_Example1(t *testing.T) {
	changes := []int{3, 3, 4, -2, -4}
	
	result := CalibrateFreq(0, changes)
	fmt.Printf("*** Result: %d\n", result)

	if result != 10 {
		t.Logf("Expected frequency %d, but got %d\n", 10, result)
		t.Fail()
	}
}

func TestCalibrateFreq_Example2(t *testing.T) {
	changes := []int{-6, 3, 8, 5, -6}
	
	result := CalibrateFreq(0, changes)
	fmt.Printf("*** Result: %d\n", result)

	if result != 5 {
		t.Logf("Expected frequency %d, but got %d\n", 5, result)
		t.Fail()
	}
}

func TestCalibrateFreq_Example3(t *testing.T) {
	changes := []int{7, 7, -2, -7, -4}
	
	result := CalibrateFreq(0, changes)
	fmt.Printf("*** Result: %d\n", result)

	if result != 14 {
		t.Logf("Expected frequency %d, but got %d\n", 14, result)
		t.Fail()
	}
}
