package day8

import (
	"testing"
)

var input = []int{2, 3, 0, 3, 10, 11, 12, 1, 1, 0, 1, 99, 2, 1, 1, 2,}

func Test_ParseNode(t *testing.T) {
	node, cursor := ParseNode(input)

	if node == nil {
		t.Errorf("Node is nil")
	}

	if cursor != 16 {
		t.Errorf("Expected cursor = 16, got %d", cursor)
	}
}

func Test_SumMetadata(t *testing.T) {
	node, _ := ParseNode(input)
	sum := SumMetadata(node)

	if sum != 138 {
		t.Errorf("Expected sum = 138, got %d", sum)
	}
}

func Test_SumValues(t *testing.T) {
	node, _ := ParseNode(input)
	sum := SumValues(node)

	if sum != 66 {
		t.Errorf("Expected sum = 66, got %d", sum)
	}
}
