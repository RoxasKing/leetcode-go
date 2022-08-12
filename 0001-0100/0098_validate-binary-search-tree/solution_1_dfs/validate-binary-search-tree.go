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

func isValidBST(root *TreeNode) bool {
	var dfs func(node *TreeNode) (int, int, bool)
	dfs = func(node *TreeNode) (int, int, bool) {
		max, min := node.Val, node.Val
		if node.Left != nil {
			lmax, lmin, ok := dfs(node.Left)
			if !ok || lmax >= node.Val {
				return 1<<31 - 1, -1 << 31, false
			}
			min = lmin
		}
		if node.Right != nil {
			rmax, rmin, ok := dfs(node.Right)
			if !ok || rmin <= node.Val {
				return 1<<31 - 1, -1 << 31, false
			}
			max = rmax
		}
		return max, min, true
	}
	_, _, ok := dfs(root)
	return ok
}
