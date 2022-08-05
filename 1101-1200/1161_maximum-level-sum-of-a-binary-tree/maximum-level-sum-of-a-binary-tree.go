package main

// Difficulty:
// Medium

// Tags:
// BFS

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	o, d, v := 1, 1, -1<<31
	for q := []*TreeNode{root}; len(q) > 0; d++ {
		x, n := 0, len(q)
		for _, t := range q {
			x += t.Val
			if t.Left != nil {
				q = append(q, t.Left)
			}
			if t.Right != nil {
				q = append(q, t.Right)
			}
		}
		if v < x {
			o, v = d, x
		}
		q = q[n:]
	}
	return o
}
