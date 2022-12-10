package main

// Difficulty:
// Medium

// Tags:
// DFS
// Hash

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxProduct(root *TreeNode) int {
	o, h := 0, map[*TreeNode]int{}
	var subTreeSum func(t *TreeNode) int
	subTreeSum = func(t *TreeNode) int {
		if t == nil {
			return 0
		}
		if v, ok := h[t]; ok {
			return v
		}
		h[t] = t.Val + subTreeSum(t.Left) + subTreeSum(t.Right)
		return h[t]
	}
	var dfs func(t *TreeNode, sum int)
	dfs = func(t *TreeNode, sum int) {
		if t == nil {
			return
		}
		sum += t.Val
		l := subTreeSum(t.Left)
		r := subTreeSum(t.Right)
		o = max(o, max((sum+l)*r, l*(sum+r)))
		dfs(t.Left, sum+r)
		dfs(t.Right, sum+l)
	}
	dfs(root, 0)
	return o % (1e9 + 7)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
