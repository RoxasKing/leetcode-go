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

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	t := root
	for t != nil {
		if t.Left != nil {
			q := t.Left
			for ; q.Right != nil && q.Right != t; q = q.Right {
			}
			if q.Right != t {
				q.Right = t
				t = t.Left
				continue
			}
			q.Right = nil
		}
		if t == p {
			p = nil
		} else if p == nil {
			break
		}
		t = t.Right
	}
	return t
}
