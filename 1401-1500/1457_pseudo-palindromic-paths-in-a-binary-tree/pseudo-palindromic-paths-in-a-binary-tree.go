package main

// Difficulty:
// Medium

// Tags:
// Backtracking
// Counting

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pseudoPalindromicPaths(root *TreeNode) int {
	o, odd := 0, 0
	freq := [10]int{}
	var backtrack func(node *TreeNode)
	backtrack = func(node *TreeNode) {
		if freq[node.Val]++; freq[node.Val]&1 == 1 {
			odd++
		} else {
			odd--
		}
		if node.Left == nil && node.Right == nil && odd <= 1 {
			o++
		}
		if node.Left != nil {
			backtrack(node.Left)
		}
		if node.Right != nil {
			backtrack(node.Right)
		}
		if freq[node.Val]--; freq[node.Val]&1 == 1 {
			odd++
		} else {
			odd--
		}
	}
	backtrack(root)
	return o
}
