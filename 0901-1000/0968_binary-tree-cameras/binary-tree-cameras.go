package main

// Tags:
// DFS

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minCameraCover(root *TreeNode) int {
	_, out, _ := dfs(root)
	return out
}

func dfs(root *TreeNode) (a, b, c int) {
	if root == nil {
		return 1000, 0, 0
	}
	la, lb, lc := dfs(root.Left)
	ra, rb, rc := dfs(root.Right)
	a = lc + rc + 1               // root must install camera
	b = Min(a, Min(la+rb, ra+lb)) // minimum number of cameras needed contain root
	c = Min(a, lb+rb)             // minimum number of cameras needed except root
	return
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
