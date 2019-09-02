package main

import "fmt"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return is_symmetric(root.Left, root.Right)
}

func is_symmetric(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil || right == nil {
		return false
	} else {
		if left.Val != right.Val {
			return false
		}
		return is_symmetric(left.Left, right.Right) && is_symmetric(left.Right, right.Left)
	}
}

func main() {
	head := &TreeNode{Val: 1}
	left1 := &TreeNode{Val: 2}
	right1 := &TreeNode{Val: 2}
	head.Left = left1
	head.Right = right1
	left2 := &TreeNode{Val: 3}
	//right2 := &TreeNode{Val: 4}
	left1.Left = left2
	//left1.Right = right2
	left3 := &TreeNode{Val: 4}
	//right3 := &TreeNode{Val: 3}
	right1.Left = left3
	//right1.Right = right3

	fmt.Println(isSymmetric(head))
}
