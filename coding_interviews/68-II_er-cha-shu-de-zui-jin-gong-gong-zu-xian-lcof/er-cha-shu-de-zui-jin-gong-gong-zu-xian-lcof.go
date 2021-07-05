package main

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return dfs(root, p, q)
}

func dfs(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	} else if root == p || root == q {
		return root
	}
	l := dfs(root.Left, p, q)
	r := dfs(root.Right, p, q)
	if l != nil && r != nil {
		return root
	} else if l != nil {
		return l
	}
	return r
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
