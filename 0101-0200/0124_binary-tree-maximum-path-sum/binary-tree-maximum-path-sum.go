package main

// Tags:
// DFS

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	out := -1 << 31
	dfs(root, &out)
	return out
}

func dfs(root *TreeNode, out *int) int {
	if root == nil {
		return 0
	}
	lSum := Max(dfs(root.Left, out), 0)
	rSum := Max(dfs(root.Right, out), 0)
	*out = Max(*out, root.Val+lSum+rSum)
	return root.Val + Max(lSum, rSum)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
