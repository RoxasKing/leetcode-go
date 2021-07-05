package main

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var out [][]int
	q := []*TreeNode{root}
	for len(q) != 0 {
		size := len(q)
		tmp := make([]int, size)
		for i := range q {
			tmp[i] = q[i].Val
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}
		out = append(out, tmp)
		q = q[size:]
	}
	for i := 0; i < len(out)>>1; i++ {
		out[i], out[len(out)-1-i] = out[len(out)-1-i], out[i]
	}
	return out
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
