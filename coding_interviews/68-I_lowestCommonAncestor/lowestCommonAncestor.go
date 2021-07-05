package main

// Tags:
// DFS + Recursion

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val > q.Val {
		p, q = q, p
	}
	return dfs(root, p, q)
}

func dfs(root, p, q *TreeNode) *TreeNode {
	if root.Val < p.Val {
		return dfs(root.Right, p, q)
	} else if root.Val > q.Val {
		return dfs(root.Left, p, q)
	}
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
