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

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	var parent *TreeNode
	for t := root; t != nil; parent, t = t, t.Right {
		if t.Val < val {
			if parent == nil {
				return &TreeNode{val, t, nil}
			}
			parent.Right = &TreeNode{val, t, nil}
			return root
		}
	}
	parent.Right = &TreeNode{val, nil, nil}
	return root
}
