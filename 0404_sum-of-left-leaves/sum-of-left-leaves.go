package main

/*
  计算给定二叉树的所有左叶子之和。
*/

// DFS + Recursion
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var sum int
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		sum += root.Left.Val
	}
	sum += sumOfLeftLeaves(root.Left)
	sum += sumOfLeftLeaves(root.Right)
	return sum
}

// BFS
func sumOfLeftLeaves2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var sum int
	q := []*TreeNode{root}
	for len(q) != 0 {
		n := q[0]
		q = q[1:]
		if n.Left != nil {
			if n.Left.Left == nil && n.Left.Right == nil {
				sum += n.Left.Val
			}
			q = append(q, n.Left)
		}
		if n.Right != nil {
			q = append(q, n.Right)
		}
	}
	return sum
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
