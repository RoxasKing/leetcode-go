package main

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 != nil && t2 != nil {
		return &TreeNode{
			Val:   t1.Val + t2.Val,
			Left:  mergeTrees(t1.Left, t2.Left),
			Right: mergeTrees(t1.Right, t2.Right),
		}
	} else if t1 != nil {
		return t1
	} else if t2 != nil {
		return t2
	}
	return nil
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
