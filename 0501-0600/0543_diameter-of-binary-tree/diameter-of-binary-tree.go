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
func diameterOfBinaryTree(root *TreeNode) int {
	out := 0
	dfs(root, 0, &out)
	return out
}

func dfs(root *TreeNode, depth int, out *int) int {
	if root == nil {
		return depth
	}

	depth++
	ld := dfs(root.Left, depth, out)
	rd := dfs(root.Right, depth, out)
	*out = Max(*out, ld+rd-depth<<1)
	return Max(ld, rd)
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
