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

func recoverTree(root *TreeNode) {
	var a, b, p *TreeNode
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
		if p != nil && p.Val > t.Val {
			if a == nil {
				a = p
			}
			b = t
		}
		p = t
		t = t.Right
	}
	a.Val, b.Val = b.Val, a.Val
}
