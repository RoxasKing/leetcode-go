package main

// Difficulty:
// Hard

// Tags:
// DFS

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minCameraCover(root *TreeNode) int {
	var dfs func(t *TreeNode) (a, b, c int)
	dfs = func(t *TreeNode) (a int, b int, c int) {
		if t == nil {
			return 1000, 0, 0
			// Attention: The return of 'a' must be maximized in case its parent and grandparent don't install a camera.
		}
		la, lb, lc := dfs(t.Left)
		ra, rb, rc := dfs(t.Right)
		a = lc + rc + 1               // current TreeNode install a camera
		b = min(a, min(la+rb, ra+lb)) // current TreeNode's children install cameras
		c = min(a, lb+rb)             // current TreeNode's children don't install cameras
		return
	}
	_, o, _ := dfs(root)
	return o
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
