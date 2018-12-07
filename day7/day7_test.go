package day7

import (
	"testing"
)

var input = []string{
	"Step C must be finished before step A can begin.",
	"Step C must be finished before step F can begin.",
	"Step A must be finished before step B can begin.",
	"Step A must be finished before step D can begin.",
	"Step B must be finished before step E can begin.",
	"Step D must be finished before step E can begin.",
	"Step F must be finished before step E can begin.",
}

func TestOrderSteps1Worker(t *testing.T) {
	steps, seconds := OrderSteps(input, 1, 0)

	if steps != "CABDFE" {
		t.Logf("Expected 'CABDFE', but got %s", steps)
		t.Fail()
	}

	if seconds != 21 {
		t.Logf("Expected 21, but got %d", seconds)
		t.Fail()
	}
}

func TestOrderSteps2Workers(t *testing.T) {
	steps, seconds := OrderSteps(input, 2, 0)

	if steps != "CABFDE" {
		t.Logf("Expected 'CABFDE', but got %s", steps)
		t.Fail()
	}

	if seconds != 15 {
		t.Logf("Expected 15, but got %d", seconds)
		t.Fail()
	}
}
