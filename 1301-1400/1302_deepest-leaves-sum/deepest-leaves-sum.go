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

func deepestLeavesSum(root *TreeNode) int {
	o := 0
	for q := []*TreeNode{root}; len(q) > 0; {
		n := len(q)
		x := 0
		for _, t := range q {
			x += t.Val
			if t.Left != nil {
				q = append(q, t.Left)
			}
			if t.Right != nil {
				q = append(q, t.Right)
			}
		}
		q = q[n:]
		o = x
	}
	return o
}
