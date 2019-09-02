package main

import (
	"fmt"
	"math"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return isValid(root, math.MinInt64, math.MaxInt64)
}

func isValid(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}

	return isValid(root.Left, min, root.Val) && isValid(root.Right, root.Val, max)
}

func main() {
	head := &TreeNode{Val: 2}
	left := &TreeNode{Val: 1}
	right := &TreeNode{Val: 3}
	head.Left = left
	head.Right = right
	fmt.Println(isValidBST(head))
}
