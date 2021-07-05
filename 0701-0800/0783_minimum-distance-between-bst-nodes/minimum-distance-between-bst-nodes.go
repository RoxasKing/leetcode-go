package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDiffInBST(root *TreeNode) int {
	out := 1<<31 - 1
	preVal := -1
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
			pre.Right = nil
		}
		if preVal != -1 && root.Val-preVal < out {
			out = root.Val - preVal
		}
		preVal = root.Val
		root = root.Right
	}
	return out
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
