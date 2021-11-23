package main

// Difficulty:
// Easy

// Tags:
// DFS

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	_, out := dfs(root)
	return out
}

func dfs(node *TreeNode) (int, int) {
	if node == nil {
		return 0, 0
	}
	ls, lt := dfs(node.Left)
	rs, rt := dfs(node.Right)
	return node.Val + ls + rs, Abs(ls-rs) + lt + rt
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
