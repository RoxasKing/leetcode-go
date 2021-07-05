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
func kthLargest(root *TreeNode, k int) int {
	for root != nil {
		if root.Right != nil {
			pre := root.Right
			for pre.Left != nil && pre.Left != root {
				pre = pre.Left
			}
			if pre.Left != root {
				pre.Left = root
				root = root.Right
				continue
			}
			pre.Left = nil
		}
		k--
		if k == 0 {
			break
		}
		root = root.Left
	}
	return root.Val
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
