package main

// Difficulty:
// Easy

// Tags:
// BFS

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	o := []float64{}
	for q := []*TreeNode{root}; len(q) > 0; {
		n := len(q)
		sum := 0
		for _, t := range q {
			sum += t.Val
			if t.Left != nil {
				q = append(q, t.Left)
			}
			if t.Right != nil {
				q = append(q, t.Right)
			}
		}
		o = append(o, float64(sum)/float64(n))
		q = q[n:]
	}
	return o
}
