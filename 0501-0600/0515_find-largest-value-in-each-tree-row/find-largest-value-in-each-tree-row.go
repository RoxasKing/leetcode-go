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

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	o := []int{}
	for q := []*TreeNode{root}; len(q) > 0; {
		n := len(q)
		max := -1 << 31
		for _, t := range q {
			if max < t.Val {
				max = t.Val
			}
			if t.Left != nil {
				q = append(q, t.Left)
			}
			if t.Right != nil {
				q = append(q, t.Right)
			}
		}
		o = append(o, max)
		q = q[n:]
	}
	return o
}
