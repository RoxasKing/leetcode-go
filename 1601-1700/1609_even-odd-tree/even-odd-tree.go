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

func isEvenOddTree(root *TreeNode) bool {
	q := []*TreeNode{root}
	flg := true
	for len(q) > 0 {
		n := len(q)
		for i := 0; i < n; i++ {
			if flg && (q[i].Val&1 == 0 || i > 0 && q[i].Val <= q[i-1].Val) || !flg && (q[i].Val&1 == 1 || i > 0 && q[i].Val >= q[i-1].Val) {
				return false
			}
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}
		q = q[n:]
		flg = !flg
	}
	return true
}
