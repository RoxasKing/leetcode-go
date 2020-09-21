package main

/*
  给定一个二叉搜索树（Binary Search Tree），把它转换成为累加树（Greater Tree)，使得每个节点的值是原来的节点值加上所有大于它的节点值之和。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/convert-bst-to-greater-tree
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

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
