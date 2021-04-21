package main

/*
  翻转一棵二叉树。
*/

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	_ = invertTree(root.Left)
	_ = invertTree(root.Right)
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
