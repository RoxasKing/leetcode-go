package main

import "strconv"

// Difficulty:
// Easy

// Tags:
// DFS

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(root *TreeNode) string {
	if root == nil {
		return ""
	}
	out := strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		return out
	}
	out += "(" + tree2str(root.Left) + ")"
	if root.Right != nil {
		out += "(" + tree2str(root.Right) + ")"
	}
	return out
}
