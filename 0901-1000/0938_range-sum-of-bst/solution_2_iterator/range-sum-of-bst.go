package main

// Difficulty:
// Easy

// Tags:
// Iterator

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	out := 0
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
		if low <= root.Val && root.Val <= high {
			out += root.Val
		}
		root = root.Right
	}
	return out
}
