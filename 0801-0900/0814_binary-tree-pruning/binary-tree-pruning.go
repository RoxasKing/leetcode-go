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

func pruneTree(root *TreeNode) *TreeNode {
	var dfs func(node *TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		ok := node.Val == 1
		if !dfs(node.Left) {
			node.Left = nil
		} else {
			ok = true
		}
		if !dfs(node.Right) {
			node.Right = nil
		} else {
			ok = true
		}
		return ok
	}
	if !dfs(root) {
		return nil
	}
	return root
}
