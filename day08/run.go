package day08

import "fmt"

func Run() {
	fmt.Println("Running Day 08 solution...")
	// Call your solution here

	root := &Node{
		Value: 1,
		Left: &Node{
			Value: 0,
			Left:  &Node{Value: 1},
			Right: &Node{Value: 0},
		},
		Right: &Node{
			Value: 1,
			Left: &Node{Value: 1,
				Left:  &Node{Value: 1},
				Right: &Node{Value: 0}},
			Right: &Node{Value: 0},
		},
	}
	PrintTree(root, 0)
	fmt.Println("Count of unival subtrees:", countUnivalSubtreesOptimized(root))
}
