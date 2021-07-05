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
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	out := [][]int{}
	q := []*TreeNode{root}
	for len(q) > 0 {
		size := len(q)
		tmp := []int{}
		for _, t := range q {
			tmp = append(tmp, t.Val)
			if t.Left != nil {
				q = append(q, t.Left)
			}
			if t.Right != nil {
				q = append(q, t.Right)
			}
		}
		out = append(out, tmp)
		q = q[size:]
	}
	return out
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
