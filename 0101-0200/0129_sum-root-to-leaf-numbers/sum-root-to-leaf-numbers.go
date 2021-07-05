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
func sumNumbers(root *TreeNode) int {
	out := 0
	dfs(root, root.Val, &out)
	return out
}

func dfs(root *TreeNode, cur int, out *int) {
	if root.Left == nil && root.Right == nil {
		*out += cur
	}

	if root.Left != nil {
		dfs(root.Left, cur*10+root.Left.Val, out)
	}

	if root.Right != nil {
		dfs(root.Right, cur*10+root.Right.Val, out)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
