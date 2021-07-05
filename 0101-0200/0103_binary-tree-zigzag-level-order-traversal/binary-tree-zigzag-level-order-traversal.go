package main

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	out := [][]int{}
	queue := []*TreeNode{root}
	reverse := false
	for len(queue) > 0 {
		n := len(queue)
		tmp := make([]int, n)
		for i := 0; i < n; i++ {
			tmp[i] = queue[i].Val
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		if reverse {
			for i := 0; i < n>>1; i++ {
				tmp[i], tmp[n-1-i] = tmp[n-1-i], tmp[i]
			}
		}
		out = append(out, tmp)
		queue = queue[n:]
		reverse = !reverse
	}
	return out
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
