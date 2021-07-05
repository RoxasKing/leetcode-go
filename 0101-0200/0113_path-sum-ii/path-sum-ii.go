package main

// Tags:
// Backtracking

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}
	out := [][]int{}
	dfs(root, targetSum-root.Val, []int{root.Val}, &out)
	return out
}

func dfs(root *TreeNode, targetSum int, cur []int, out *[][]int) {
	if root.Left == nil && root.Right == nil {
		if targetSum == 0 {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			*out = append(*out, tmp)
		}
		return
	}

	if root.Left != nil {
		targetSum -= root.Left.Val
		cur = append(cur, root.Left.Val)
		dfs(root.Left, targetSum, cur, out)
		cur = cur[:len(cur)-1]
		targetSum += root.Left.Val
	}

	if root.Right != nil {
		targetSum -= root.Right.Val
		cur = append(cur, root.Right.Val)
		dfs(root.Right, targetSum, cur, out)
		// cur = cur[:len(cur)-1]
		// targetSum += root.Right.Val
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
