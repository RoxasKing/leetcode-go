package main

// Tags:
// Dynamic Programming + DFS

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
	a, b := dfs(root)
	return Max(a, b)
}

func dfs(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	l0, l1 := dfs(root.Left)
	r0, r1 := dfs(root.Right)

	a := root.Val + l1 + r1
	b := Max(l0, l1) + Max(r0, r1)

	return a, b
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
