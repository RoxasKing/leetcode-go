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

func levelOrder(root *TreeNode) [][]int {
	o := [][]int{}
	if root == nil {
		return o
	}
	for q := []*TreeNode{root}; len(q) > 0; {
		n := len(q)
		a := make([]int, n)
		for i, t := range q {
			a[i] = t.Val
			if t.Left != nil {
				q = append(q, t.Left)
			}
			if t.Right != nil {
				q = append(q, t.Right)
			}
		}
		o = append(o, a)
		q = q[n:]
	}
	return o
}
