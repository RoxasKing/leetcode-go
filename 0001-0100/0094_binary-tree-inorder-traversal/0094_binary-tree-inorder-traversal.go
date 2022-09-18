package main

// Difficulty:
// Easy

// Tags:
// Morris Traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	o := []int{}
	for t := root; t != nil; {
		if t.Left != nil {
			p := t.Left
			for ; p.Right != nil && p.Right != t; p = p.Right {
			}
			if p.Right != t {
				p.Right = t
				t = t.Left
				continue
			}
			p.Right = nil
		}
		o = append(o, t.Val)
		t = t.Right
	}
	return o
}
