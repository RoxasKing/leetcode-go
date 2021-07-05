package main

import "sort"

// DFS
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

// Binary Search + Bit Manipulation
func countNodes2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 0
	for n := root; n.Left != nil; n = n.Left {
		depth++
	}
	return sort.Search(1<<(depth+1), func(i int) bool {
		if i <= 1<<depth {
			return false
		}
		bits := 1 << (depth - 1)
		n := root
		for n != nil && bits > 0 {
			if bits&i == 0 {
				n = n.Left
			} else {
				n = n.Right
			}
			bits >>= 1
		}
		return n == nil
	}) - 1
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
