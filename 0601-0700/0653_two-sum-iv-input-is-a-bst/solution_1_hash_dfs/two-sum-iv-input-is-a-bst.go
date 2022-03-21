package main

// Difficulty:
// Easy

// Tags:
// Hash
// DFS

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	h := map[int]struct{}{}
	var dfs func(*TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if _, ok := h[k-node.Val]; ok {
			return true
		}
		h[node.Val] = struct{}{}
		if node.Left != nil && dfs(node.Left) || node.Right != nil && dfs(node.Right) {
			return true
		}
		return false
	}
	return dfs(root)
}
