package main

// Difficulty:
// Medium

// Tags:
// Stack

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	var last, a, b *TreeNode
	node := root
	stack := []*TreeNode{}
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if last != nil && node.Val < last.Val {
			if a == nil {
				a = last
			}
			b = node
		}
		last = node
		node = node.Right
	}
	a.Val, b.Val = b.Val, a.Val
}
