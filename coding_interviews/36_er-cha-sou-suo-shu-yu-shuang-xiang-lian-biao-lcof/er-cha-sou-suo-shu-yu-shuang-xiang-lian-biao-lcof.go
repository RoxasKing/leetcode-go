package main

// Tags:
// Morris Traversal
func treeToDoublyList(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var head, last *TreeNode
	node := root
	for node != nil {
		if node.Left != nil {
			nPre := node.Left
			for nPre.Right != nil && nPre.Right != node {
				nPre = nPre.Right
			}
			if nPre.Right != node {
				nPre.Right = node
				node = node.Left
				continue
			}
		}
		if head == nil {
			head = node
		}
		if last != nil {
			last.Right = node
			node.Left = last
		}
		last = node
		node = node.Right
	}
	last.Right = head
	head.Left = last
	return head
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
