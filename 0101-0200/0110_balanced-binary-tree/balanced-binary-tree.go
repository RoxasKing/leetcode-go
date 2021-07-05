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
	_, ok := dfs(root, 0)
	return ok
}

func dfs(root *TreeNode, depth int) (int, bool) {
	if root == nil {
		return depth, true
	}
	depth++
	ld, lok := dfs(root.Left, depth)
	if !lok {
		return depth, false
	}
	rd, rok := dfs(root.Right, depth)
	if !rok {
		return depth, false
	}
	return Max(ld, rd), Abs(ld-rd) < 2
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
