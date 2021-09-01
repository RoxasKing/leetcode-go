package main

import "runtime/debug"

// Tags:
// Dynamic Programming
// DFS

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxValue(root *TreeNode, k int) int {
	_, out := dfs(root, k)
	return out
}

func dfs(t *TreeNode, k int) ([]int, int) {
	f := make([]int, k+1)
	if t == nil {
		return f, 0
	}
	lf, lmax := dfs(t.Left, k)
	rf, rmax := dfs(t.Right, k)
	f[0] = lmax + rmax
	max := f[0]
	for i := 1; i <= k; i++ {
		for j := 0; j < i; j++ {
			f[i] = Max(f[i], t.Val+lf[j]+rf[i-1-j])
		}
		max = Max(max, f[i])
	}
	return f, max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func init() { debug.SetGCPercent(-1) }
