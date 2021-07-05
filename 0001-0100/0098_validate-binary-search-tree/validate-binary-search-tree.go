package main

// Tags:
// Morris Traversal

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	var prev *TreeNode
	for root != nil {
		if root.Left != nil {
			rootPre := root.Left
			for rootPre.Right != nil && rootPre.Right != root {
				rootPre = rootPre.Right
			}
			if rootPre.Right != root {
				rootPre.Right = root
				root = root.Left
				continue
			}
			rootPre.Right = nil
		}
		if prev != nil && prev.Val >= root.Val {
			return false
		}
		prev, root = root, root.Right
	}
	return true
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
