package main

// Difficulty:
// Medium

// Tags:
// DFS

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == root.Right {
			return nil
		}
		if root.Left != nil {
			prev := root.Left
			for ; prev.Right != nil; prev = prev.Right {
			}
			root.Val = prev.Val
			root.Left = deleteNode(root.Left, root.Val)
		} else {
			next := root.Right
			for ; next.Left != nil; next = next.Left {
			}
			root.Val = next.Val
			root.Right = deleteNode(root.Right, root.Val)
		}
	}
	return root
}
