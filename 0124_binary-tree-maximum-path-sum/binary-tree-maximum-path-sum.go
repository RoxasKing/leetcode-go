package main

/*
  给定一个非空二叉树，返回其最大路径和。
  本题中，路径被定义为一条从树中任意节点出发，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。
*/

// DFS + Recursion
func maxPathSum(root *TreeNode) int {
	max := -1 << 31
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		lMax := Max(dfs(node.Left), 0)
		rMax := Max(dfs(node.Right), 0)
		max = Max(max, node.Val+lMax+rMax)
		return node.Val + Max(lMax, rMax)
	}
	_ = dfs(root)
	return max
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
