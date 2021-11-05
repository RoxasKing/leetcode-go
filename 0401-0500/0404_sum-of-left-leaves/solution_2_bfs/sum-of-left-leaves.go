package main

// Difficulty:
// Easy

// BFS

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	out := 0
	q := []*TreeNode{root}
	for ; len(q) != 0; q = q[1:] {
		node := q[0]
		if node.Left != nil {
			if node.Left.Left == nil && node.Left.Right == nil {
				out += node.Left.Val
			} else {
				q = append(q, node.Left)
			}
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
	}
	return out
}
