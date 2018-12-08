package day8

// Node is an object in a tree structure
type Node struct {
	children []*Node
	metadata []int
}

// ParseNode builds the node structure from a slice of integers
func ParseNode(input []int) (*Node, int) {
	childCount := input[0]
	metadataCount := input[1]

	n := &Node{}

	cursor := 2
	for len(n.children) < childCount {
		c, jump := ParseNode(input[cursor:])
		n.children = append(n.children, c)
		cursor += jump
	}

	n.metadata = input[cursor:cursor+metadataCount]
	cursor += metadataCount

	return n, cursor
}

// SumMetadata gets the sum of the metadata values from a node's tree
func SumMetadata(n *Node) int {
	sum := 0
	for _, v := range n.metadata {
		sum += v
	}

	for i := range n.children {
		sum += SumMetadata(n.children[i])
	}

	return sum
}

// SumValues get the sum of "node values" for all nodes in the tree
func SumValues(n *Node) int {
	sum := 0
	if len(n.children) == 0 {
		for _, v := range n.metadata {
			sum += v
		}
		return sum
	}

	for _, v := range n.metadata {
		if len(n.children) >= v {
			sum += SumValues(n.children[v-1])
		}
	}

	return sum
}
