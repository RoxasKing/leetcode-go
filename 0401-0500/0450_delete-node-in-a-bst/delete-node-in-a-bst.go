package main

// Tags:
// DFS

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == root.Right {
			return nil
		}
		if root.Left != nil {
			pre := root.Left
			for pre.Right != nil {
				pre = pre.Right
			}
			root.Val = pre.Val
			root.Left = deleteNode(root.Left, root.Val)
		} else {
			nxt := root.Right
			for nxt.Left != nil {
				nxt = nxt.Left
			}
			root.Val = nxt.Val
			root.Right = deleteNode(root.Right, root.Val)
		}
	}
	return root
}
