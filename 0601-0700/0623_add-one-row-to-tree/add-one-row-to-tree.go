package main

// Difficulty:
// Medium

// Tags:
// BFS

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		return &TreeNode{Val: val, Left: root}
	}
	q := []*TreeNode{root}
	for ; depth > 2; depth-- {
		n := len(q)
		for _, t := range q {
			if t.Left != nil {
				q = append(q, t.Left)
			}
			if t.Right != nil {
				q = append(q, t.Right)
			}
		}
		q = q[n:]
	}
	for _, t := range q {
		l, r := t.Left, t.Right
		t.Left = &TreeNode{Val: val}
		t.Right = &TreeNode{Val: val}
		if l != nil {
			t.Left.Left = l
		}
		if r != nil {
			t.Right.Right = r
		}
	}
	return root
}
