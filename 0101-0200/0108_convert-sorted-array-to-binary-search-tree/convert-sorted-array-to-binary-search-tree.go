package main

// Difficulty:
// Easy

// Tags:
// DFS

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	n := len(nums)
	var build func(l, r int) *TreeNode
	build = func(l, r int) *TreeNode {
		if l > r {
			return nil
		}
		m := (l + r + 1) / 2
		return &TreeNode{nums[m], build(l, m-1), build(m+1, r)}
	}
	return build(0, n-1)
}
