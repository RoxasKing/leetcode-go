package main

func getMinimumDifference(root *TreeNode) int {
	node := root
	min := 1<<31 - 1
	var pre *TreeNode
	for node != nil {
		if node.Left != nil {
			pre := node.Left
			for pre.Right != nil && pre.Right != node {
				pre = pre.Right
			}
			if pre.Right != node {
				pre.Right = node
				node = node.Left
				continue
			}
			pre.Right = nil
		}
		if pre != nil {
			min = Min(min, node.Val-pre.Val)
		}
		pre = node
		node = node.Right
	}
	return min
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
