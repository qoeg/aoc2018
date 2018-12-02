package day2

import (
	"testing"
)

var examples1 = []string{
	"abcdef",
	"bababc",
	"abbcde",
	"abcccd",
	"aabcdd",
	"abcdee",
	"ababab",
}

var examples2 = []string{
	"abcde",
	"fghij",
	"klmno",
	"pqrst",
	"fguij",
	"axcye",
	"wvxyz",
}

func TestCountContainsTwo1(t *testing.T) {
	count := CountContainsTwo(examples1)
	if count != 4 {
		t.Logf("Expected %d, but got %d\n", 4, count)
		t.Fail()
	}
}

func TestCheckContainsThree1(t *testing.T) {
	count := CountContainsThree(examples1)
	if count != 3 {
		t.Logf("Expected %d, but got %d\n", 3, count)
		t.Fail()
	}
}

func TestGetCorrectID(t *testing.T) {
	result := GetCorrectID(examples2)
	if result != "fgij" {
		t.Logf("Expected %s, but got %s\n", "fgij", result)
		t.Fail()
	}
}
