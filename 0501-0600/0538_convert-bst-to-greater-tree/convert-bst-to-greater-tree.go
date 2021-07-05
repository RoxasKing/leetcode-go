package main

// Tags:
// Morris Traversal
func convertBST(root *TreeNode) *TreeNode {
	node := root
	var next *TreeNode
	for node != nil {
		if node.Right != nil {
			pre := node.Right
			for pre.Left != nil && pre.Left != node {
				pre = pre.Left
			}
			if pre.Left != node {
				pre.Left = node
				node = node.Right
				continue
			}
			pre.Left = nil
		}
		if next != nil {
			node.Val += next.Val
		}
		next = node
		node = node.Left
	}
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
