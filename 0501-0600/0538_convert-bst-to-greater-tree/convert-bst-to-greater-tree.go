package main

// Difficulty:
// Medium

// Tags:
// Morris Traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	for sum, ptr := 0, root; ptr != nil; {
		if ptr.Right != nil {
			next := ptr.Right
			for ; next.Left != nil && next.Left != ptr; next = next.Left {
			}
			if next.Left != ptr {
				next.Left = ptr
				ptr = ptr.Right
				continue
			}
			next.Left = nil
		}
		ptr.Val += sum
		sum = ptr.Val
		ptr = ptr.Left
	}
	return root
}
