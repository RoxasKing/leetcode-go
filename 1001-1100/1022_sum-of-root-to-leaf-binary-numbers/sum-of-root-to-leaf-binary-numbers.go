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

func sumRootToLeaf(root *TreeNode) int {
	var out int
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, num int) {
		num = num<<1 + node.Val
		if node.Left == nil && node.Right == nil {
			out += num
		}
		if node.Left != nil {
			dfs(node.Left, num)
		}
		if node.Right != nil {
			dfs(node.Right, num)
		}
	}
	dfs(root, 0)
	return out
}
