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

func longestUnivaluePath(root *TreeNode) int {
	o := 0
	var dfs func(t *TreeNode) int
	dfs = func(t *TreeNode) int {
		if t == nil {
			return 0
		}
		cnt, ret := 0, 1
		if lc := dfs(t.Left); t.Left != nil && t.Left.Val == t.Val {
			cnt += lc
			ret = max(ret, 1+lc)
		}
		if rc := dfs(t.Right); t.Right != nil && t.Right.Val == t.Val {
			cnt += rc
			ret = max(ret, 1+rc)
		}
		o = max(o, cnt)
		return ret
	}
	dfs(root)
	return o
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
