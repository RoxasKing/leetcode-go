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

func constructMaximumBinaryTree(nums []int) *TreeNode {
	var build func(l, r int) *TreeNode
	build = func(l, r int) *TreeNode {
		if l > r {
			return nil
		}
		m := l
		for i := l + 1; i <= r; i++ {
			if nums[m] < nums[i] {
				m = i
			}
		}
		return &TreeNode{nums[m], build(l, m-1), build(m+1, r)}
	}
	return build(0, len(nums)-1)
}
