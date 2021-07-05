package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
	var head, pre *TreeNode
	for root != nil {
		if root.Left != nil {
			pre := root.Left
			for pre.Right != nil && pre.Right != root {
				pre = pre.Right
			}
			if pre.Right != root {
				pre.Right = root
				root = root.Left
				continue
			}
			root.Left = nil
		}
		if head == nil {
			head = root
		}
		if pre != nil {
			pre.Right = root
		}
		pre, root = root, root.Right
	}
	return head
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
