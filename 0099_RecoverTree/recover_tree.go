package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 1 2 3 4 5      1 4 3 2 5
func recoverTree(root *TreeNode) {
	save := make([]*TreeNode, 0)
	stack := make([]*TreeNode, 0)
	t := root
	for t != nil || len(stack) != 0 {
		for t != nil {
			stack = append(stack, t)
			t = t.Left
		}
		if len(stack) != 0 {
			if stack[len(stack)-1] != nil {
				save = append(save, stack[len(stack)-1])
			}
			t = stack[len(stack)-1].Right
			stack = stack[:len(stack)-1]
		}
	}
	a, b := -1, -1
	for i := 1; i < len(save); i++ {
		if save[i].Val < save[i-1].Val {
			a = i - 1
			break
		}
	}
	for i := len(save) - 2; i >= a; i-- {
		if save[i].Val > save[i+1].Val {
			b = i + 1
			break
		}
	}
	save[a].Val, save[b].Val = save[b].Val, save[a].Val
}

func main() {
	root := &TreeNode{Val: 3}
	left := &TreeNode{Val: 1}
	root.Left = left
	right := &TreeNode{Val: 4}
	root.Right = right
	left1 := &TreeNode{Val: 2}
	right.Left = left1

	recoverTree(root)
}
