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

func increasingBST(root *TreeNode) *TreeNode {
	var o, p *TreeNode
	for t := root; t != nil; {
		if t.Left != nil {
			pre := t.Left
			for ; pre.Right != nil && pre.Right != t; pre = pre.Right {
			}
			if pre.Right != t {
				pre.Right = t
				t = t.Left
				continue
			}
			pre.Right = nil
		}
		t.Left = nil
		if p == nil {
			o, p = t, t
		} else {
			p.Right = t
			p = p.Right
		}
		t = t.Right
	}
	return o
}
