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

func goodNodes(root *TreeNode) int {
	o := 0
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, max int) {
		if node == nil {
			return
		}
		if node.Val >= max {
			max = node.Val
			o++
		}
		dfs(node.Left, max)
		dfs(node.Right, max)
	}
	dfs(root, -1<<31)
	return o
}
