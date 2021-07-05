package main

// Tags:
// DFS

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if equal(s, t) {
		return true
	}
	if s == nil {
		return false
	}
	return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func equal(s, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	} else if s == nil || t == nil || s.Val != t.Val {
		return false
	}
	return equal(s.Left, t.Left) && equal(s.Right, t.Right)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
