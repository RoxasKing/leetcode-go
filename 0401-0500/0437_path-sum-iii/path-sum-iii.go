package main

// Tags:
// Backtracking

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	return dfs(root, targetSum) + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)
}

func dfs(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	targetSum -= root.Val
	var out int
	if targetSum == 0 {
		out++
	}
	out += dfs(root.Left, targetSum)
	out += dfs(root.Right, targetSum)
	return out
}
