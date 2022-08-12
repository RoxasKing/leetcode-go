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

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val > q.Val {
		p, q = q, p
	}
	var dfs func(node *TreeNode, l, r int) *TreeNode
	dfs = func(node *TreeNode, l, r int) *TreeNode {
		if node.Val > r {
			return dfs(node.Left, l, r)
		} else if node.Val < l {
			return dfs(node.Right, l, r)
		}
		return node
	}
	return dfs(root, p.Val, q.Val)
}
