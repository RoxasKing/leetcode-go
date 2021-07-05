package main

// Tags:
// DFS + Recursion

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	out := []int{}
	dfs(root, 0, &out)
	return out
}

func dfs(root *TreeNode, depth int, out *[]int) {
	if root == nil {
		return
	}
	if depth > len(*out)-1 {
		*out = append(*out, root.Val)
	} else {
		(*out)[depth] = root.Val
	}
	dfs(root.Left, depth+1, out)
	dfs(root.Right, depth+1, out)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
