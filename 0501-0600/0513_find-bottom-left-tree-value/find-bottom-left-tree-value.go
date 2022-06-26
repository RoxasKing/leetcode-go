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

func findBottomLeftValue(root *TreeNode) int {
	o, d := root.Val, 1
	var dfs func(node *TreeNode, deep int)
	dfs = func(node *TreeNode, deep int) {
		deep++
		if deep > d {
			o, d = node.Val, deep
		}
		if node.Left != nil {
			dfs(node.Left, deep)
		}
		if node.Right != nil {
			dfs(node.Right, deep)
		}
	}
	dfs(root, 0)
	return o
}
