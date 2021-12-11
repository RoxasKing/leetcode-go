package main

// Difficulty:
// Medium

// Tags:
// DFS
// Dynamic Programming

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	var dfs func(*TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		la, lb := dfs(node.Left)
		ra, rb := dfs(node.Right)
		b := la + ra
		a := Max(b, node.Val+lb+rb)
		return a, b // a: with root node ; b: without root node
	}
	out, _ := dfs(root)
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
