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
	var dfs func(node *TreeNode) (*TreeNode, bool)
	dfs = func(node *TreeNode) (*TreeNode, bool) {
		if node == nil {
			return nil, false
		}
		lt, lb := dfs(node.Left)
		rt, rb := dfs(node.Right)
		if node.Val == 0 && !lb && !rb {
			return nil, false
		}
		node.Left = lt
		node.Right = rt
		return node, true
	}
	o, _ := dfs(root)
	return o
}
