package main

import "fmt"

// Difficulty:
// Medium

// Tags:
// DFS
// Hash

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	h := map[string]int{}
	o := []*TreeNode{}
	var dfs func(node *TreeNode) string
	dfs = func(node *TreeNode) string {
		if node == nil {
			return ""
		}
		s := fmt.Sprintf("%d(%s)(%s)", node.Val, dfs(node.Left), dfs(node.Right))
		if c, ok := h[s]; ok && c == 1 {
			o = append(o, node)
		}
		h[s]++
		return s
	}
	dfs(root)
	return o
}
