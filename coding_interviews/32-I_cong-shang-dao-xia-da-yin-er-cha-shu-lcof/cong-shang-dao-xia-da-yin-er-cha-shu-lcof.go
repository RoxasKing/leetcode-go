package main

// Tags:
// Preorder traversal (VLR) + BFS

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	q := []*TreeNode{root}
	out := []int{}
	for len(q) > 0 {
		e := q[0]
		q = q[1:]
		out = append(out, e.Val)
		if e.Left != nil {
			q = append(q, e.Left)
		}
		if e.Right != nil {
			q = append(q, e.Right)
		}
	}
	return out
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
