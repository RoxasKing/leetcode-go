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

func flatten(root *TreeNode) {
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
			p.Right = t.Right
			t.Right = t.Left
			t.Left = nil
		}
		t = t.Right
	}
}
