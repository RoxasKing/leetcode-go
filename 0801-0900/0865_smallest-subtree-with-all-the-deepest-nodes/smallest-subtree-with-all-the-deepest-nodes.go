package main

// Tags:
// DFS

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	node, _ := dfs(root, 0)
	return node
}

func dfs(root *TreeNode, depth int) (*TreeNode, int) {
	if root.Left == nil && root.Right == nil {
		return root, depth
	}
	depth++
	var tl, tr *TreeNode
	var dl, dr int
	if root.Left != nil {
		tl, dl = dfs(root.Left, depth)
	}
	if root.Right != nil {
		tr, dr = dfs(root.Right, depth)
	}
	if dl > dr {
		return tl, dl
	} else if dl < dr {
		return tr, dr
	}
	return root, dl
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
