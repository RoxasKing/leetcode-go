package main

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
	q := []*TreeNode{root}
	rev := false
	out := [][]int{}
	for len(q) > 0 {
		size := len(q)
		tmp := make([]int, 0, size)
		for _, e := range q {
			tmp = append(tmp, e.Val)
			if e.Left != nil {
				q = append(q, e.Left)
			}
			if e.Right != nil {
				q = append(q, e.Right)
			}
		}
		if rev {
			for i := 0; i < size>>1; i++ {
				tmp[i], tmp[size-1-i] = tmp[size-1-i], tmp[i]
			}
		}
		q = q[size:]
		rev = !rev
		out = append(out, tmp)
	}
	return out
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
