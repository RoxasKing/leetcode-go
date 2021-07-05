package main

// Tags:
// DFS + Dynamic Programming

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxValue(root *TreeNode, k int) int {
	dp := make(map[*TreeNode][]int)
	return dfs(dp, root, k, k)
}

func dfs(dp map[*TreeNode][]int, root *TreeNode, remain, k int) int {
	if root == nil {
		return 0
	}

	if dp[root] == nil {
		dp[root] = make([]int, k+1)
	}

	if dp[root][remain] > 0 {
		return dp[root][remain]
	}

	out := dfs(dp, root.Left, k, k) + dfs(dp, root.Right, k, k) // root 不染色

	for i := 0; i <= remain-1; i++ {
		l1 := dfs(dp, root.Left, i, k)           // left 染色 i
		r1 := dfs(dp, root.Right, remain-1-i, k) // right 染色 remain-i
		out = Max(out, root.Val+l1+r1)           // root 染色
	}
	dp[root][remain] = out
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
