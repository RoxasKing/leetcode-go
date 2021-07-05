package main

// Tags:
// BFS
func averageOfLevels(root *TreeNode) []float64 {
	var out []float64
	q := []*TreeNode{root}
	for len(q) != 0 {
		curSize := len(q)
		var sum float64
		for _, n := range q {
			sum += float64(n.Val)
			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
		}
		out = append(out, sum/float64(curSize))
		q = q[curSize:]
	}
	return out
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
