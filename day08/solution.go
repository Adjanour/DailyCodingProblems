package day08

import (
	"fmt"
	"strings"
)

// Node represents a node in a binary tree.
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// NewNode creates a new node with the given value.
func NewNode(value int) *Node {
	return &Node{Value: value}
}

// insert inserts a value into a binary search tree (BST).
// NOTE: This builds a BST, not a generic binary tree with a fixed structure.
func insert(root *Node, value int) {
	if root == nil {
		return // no insertion happens if root is nil — external logic must ensure root is initialized
	}
	if value < root.Value {
		if root.Left == nil {
			root.Left = NewNode(value)
		} else {
			insert(root.Left, value)
		}
	} else if value > root.Value {
		if root.Right == nil {
			root.Right = NewNode(value)
		} else {
			insert(root.Right, value)
		}
	}
}

// buildTree builds a binary search tree from a slice of values.
// To preserve structure from a level-order array, you'd need a different method.
func buildTree(values []int) *Node {
	if len(values) == 0 {
		return nil
	}
	root := NewNode(values[0])
	for _, value := range values[1:] {
		insert(root, value)
	}
	return root
}

// isUnivalTree checks if a subtree is univalued (all nodes have the same value).
// This is a naive version used to demonstrate correctness, not efficiency.
func isUnivalTree(root *Node) bool {
	if root == nil {
		return true
	}
	if root.Left != nil && root.Left.Value != root.Value {
		return false
	}
	if root.Right != nil && root.Right.Value != root.Value {
		return false
	}
	return isUnivalTree(root.Left) && isUnivalTree(root.Right)
}

// countUnivalSubtrees counts all unival subtrees in the binary tree.
// This version uses the naive isUnivalTree() and is O(n²) in worst-case.
func countUnivalSubtrees(root *Node) int {
	if root == nil {
		return 0
	}
	count := 0
	if isUnivalTree(root) {
		count = 1
	}
	return count + countUnivalSubtrees(root.Left) + countUnivalSubtrees(root.Right)
}

// countUnivalSubtreesOptimized counts unival subtrees using a single-pass bottom-up approach.
func countUnivalSubtreesOptimized(root *Node) int {
	count := 0
	isUnivalHelper(root, &count)
	return count
}

// isUnivalHelper recursively determines if the current subtree is unival.
// It increments the count if the subtree rooted at this node is unival.
func isUnivalHelper(node *Node, count *int) bool {
	// A nil node (child of a leaf) is considered unival (by definition)
	if node == nil {
		return true
	}

	// Check if left and right subtrees are unival
	leftUnival := isUnivalHelper(node.Left, count)
	rightUnival := isUnivalHelper(node.Right, count)

	// If either subtree is not unival, this cannot be unival
	if !leftUnival || !rightUnival {
		return false
	}

	// If left child exists and its value differs, not unival
	if node.Left != nil && node.Left.Value != node.Value {
		return false
	}

	// If right child exists and its value differs, not unival
	if node.Right != nil && node.Right.Value != node.Value {
		return false
	}

	// All conditions passed: subtree is unival
	*count++
	return true
}

// PrintTree prints the structure of the tree in a rotated format for visualization.
func PrintTree(node *Node, level int) {
	if node == nil {
		return
	}
	PrintTree(node.Right, level+1)
	fmt.Printf("%s%d\n", strings.Repeat("    ", level), node.Value)
	PrintTree(node.Left, level+1)
}
