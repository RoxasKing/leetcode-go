package main

// Difficulty:
// Easy

// Tags:
// DFS

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	d := 0
	for _, child := range root.Children {
		d = Max(d, maxDepth(child))
	}
	return 1 + d
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
