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

func maxAncestorDiff(root *TreeNode) int {
	out := -1 << 31
	var dfs func(*TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		min, max := node.Val, node.Val
		if node.Left != nil {
			lmin, lmax := dfs(node.Left)
			out = Max(out, Max(Abs(node.Val-lmin), Abs(node.Val-lmax)))
			min, max = Min(min, lmin), Max(max, lmax)
		}
		if node.Right != nil {
			rmin, rmax := dfs(node.Right)
			out = Max(out, Max(Abs(node.Val-rmin), Abs(node.Val-rmax)))
			min, max = Min(min, rmin), Max(max, rmax)
		}
		return min, max
	}
	_, _ = dfs(root)
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
