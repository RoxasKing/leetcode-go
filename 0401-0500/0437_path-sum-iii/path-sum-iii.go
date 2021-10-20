package main

// Difficulty:
// Medium

// Tags:
// DFS

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
	x := 0
	if targetSum -= root.Val; targetSum == 0 {
		x++
	}
	return x + dfs(root.Left, targetSum) + dfs(root.Right, targetSum)
}
