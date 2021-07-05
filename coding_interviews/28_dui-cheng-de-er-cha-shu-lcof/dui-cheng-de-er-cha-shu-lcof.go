package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isEqual(root.Left, root.Right)
}

func isEqual(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	} else if l == nil || r == nil || l.Val != r.Val {
		return false
	}
	return isEqual(l.Left, r.Right) && isEqual(l.Right, r.Left)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
