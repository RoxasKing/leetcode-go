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
func isCousins(root *TreeNode, x int, y int) bool {
	q := []*TreeNode{root}
	for len(q) > 0 {
		size := len(q)
		tmp := []*TreeNode{}
		for _, t := range q {
			if t.Left != nil {
				if t.Left.Val == x || t.Left.Val == y {
					tmp = append(tmp, t)
				}
				q = append(q, t.Left)
			}
			if t.Right != nil {
				if t.Right.Val == x || t.Right.Val == y {
					tmp = append(tmp, t)
				}
				q = append(q, t.Right)
			}
		}
		if len(tmp) == 2 {
			return tmp[0] != tmp[1]
		} else if len(tmp) == 1 {
			return false
		}
		q = q[size:]
	}
	return false
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
