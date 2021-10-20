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

func kthSmallest(root *TreeNode, k int) int {
	for p := root; p != nil; {
		if p.Left != nil {
			pp := p.Left
			for pp.Right != nil && pp.Right != p {
				pp = pp.Right
			}
			if pp.Right != p {
				pp.Right = p
				p = p.Left
				continue
			}
			pp.Right = nil
		}
		if k--; k == 0 {
			return p.Val
		}
		p = p.Right
	}
	return root.Val
}
