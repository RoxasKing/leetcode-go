package main

// Difficulty:
// Easy

// Tags:
// Backtracking

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	var backtrack func(node *TreeNode) bool
	backtrack = func(node *TreeNode) bool {
		targetSum -= node.Val
		defer func() { targetSum += node.Val }()
		return node.Left == nil && node.Right == nil && targetSum == 0 ||
			node.Left != nil && backtrack(node.Left) || node.Right != nil && backtrack(node.Right)
	}
	return backtrack(root)
}
