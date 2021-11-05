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

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return dfs(root, 0)
}

func dfs(node *TreeNode, cur int) int {
	cur = cur*10 + node.Val
	if node.Left == nil && node.Right == nil {
		return cur
	}
	out := 0
	if node.Left != nil {
		out += dfs(node.Left, cur)
	}
	if node.Right != nil {
		out += dfs(node.Right, cur)
	}
	return out
}
