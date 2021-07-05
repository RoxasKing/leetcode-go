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
	return equal(root.Left, root.Right)
}

func equal(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	} else if t1 == nil || t2 == nil || t1.Val != t2.Val {
		return false
	}
	return equal(t1.Left, t2.Right) && equal(t1.Right, t2.Left)
}

func isSymmetric2(root *TreeNode) bool {
	q := []*TreeNode{root}
	for len(q) > 0 {
		n := len(q)
		for i := 0; i < n>>1; i++ {
			if q[i] == nil && q[n-1-i] != nil ||
				q[i] != nil && q[n-1-i] == nil ||
				q[i] != nil && q[n-1-i] != nil && q[i].Val != q[n-1-i].Val {
				return false
			}
		}
		for _, t := range q {
			if t == nil {
				continue
			}
			q = append(q, t.Left, t.Right)
		}
		q = q[n:]
	}
	return true
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
