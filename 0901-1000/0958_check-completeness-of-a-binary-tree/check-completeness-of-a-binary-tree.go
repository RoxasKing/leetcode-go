package main

// Tags:
// BFS

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCompleteTree(root *TreeNode) bool {
	q := []*TreeNode{root}
	flg := false
	for len(q) > 0 {
		size := len(q)
		for _, t := range q {
			if t == nil {
				flg = true
				continue
			}
			if flg {
				return false
			}
			q = append(q, t.Left, t.Right)
		}
		q = q[size:]
	}
	return true
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
