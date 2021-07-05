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
func isBalanced(root *TreeNode) bool {
	_, ok := dfs(root)
	return ok
}

func dfs(root *TreeNode) (depth int, ok bool) {
	if root == nil {
		return 0, true
	}
	dL, okL := dfs(root.Left)
	if !okL {
		return 0, false
	}
	dR, okR := dfs(root.Right)
	if !okR {
		return 0, false
	}
	return Max(dL, dR) + 1, Abs(dL-dR) < 2
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
