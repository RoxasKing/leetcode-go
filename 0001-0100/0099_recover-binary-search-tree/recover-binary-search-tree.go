package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Morris Traversal
func recoverTree(root *TreeNode) {
	var last, a, b *TreeNode
	node := root
	for node != nil {
		if node.Left != nil {
			prev := node.Left
			for prev.Right != nil && prev.Right != node {
				prev = prev.Right
			}
			if prev.Right != node {
				prev.Right = node
				node = node.Left
				continue
			}
			prev.Right = nil
		}
		if last != nil && node.Val < last.Val {
			if a == nil {
				a = last
			}
			b = node
		}
		last = node
		node = node.Right
	}
	a.Val, b.Val = b.Val, a.Val
}

// Stack
func recoverTree2(root *TreeNode) {
	var last, a, b *TreeNode
	node := root
	stack := []*TreeNode{}
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if last != nil && node.Val < last.Val {
			if a == nil {
				a = last
			}
			b = node
		}
		last = node
		node = node.Right
	}
	a.Val, b.Val = b.Val, a.Val
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
