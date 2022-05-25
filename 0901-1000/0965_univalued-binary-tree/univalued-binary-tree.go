package main

// Difficulty:
// Easy

// Tags:
// DFS

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	if root.Left != nil && (root.Left.Val != root.Val || !isUnivalTree(root.Left)) ||
		root.Right != nil && (root.Right.Val != root.Val || !isUnivalTree(root.Right)) {
		return false
	}
	return true
}
