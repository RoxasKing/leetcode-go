package main

import "fmt"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p != nil && q != nil && p.Val != q.Val {
		return false
	} else if p == nil && q != nil || p != nil && q == nil {
		return false
	} else if p == nil && q == nil {
		return true
	}
	left := isSameTree(p.Left, q.Left)
	right := isSameTree(p.Right, q.Right)

	if left && right {
		return true
	}
	return false
}

func main() {
	head1 := &TreeNode{Val: 1}
	left1 := &TreeNode{Val: 2}
	//right1 := &TreeNode{Val: 3}
	head1.Left = left1
	//head1.Right = right1

	head2 := &TreeNode{Val: 1}
	//left2 := &TreeNode{Val: 2}
	right2 := &TreeNode{Val: 3}
	//head2.Left = left2
	head2.Right = right2

	fmt.Println(isSameTree(head1, head2))
}
