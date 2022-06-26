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

func findFrequentTreeSum(root *TreeNode) []int {
	h := map[int]int{}
	a, x := []int{}, 0
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		o := node.Val + dfs(node.Left) + dfs(node.Right)
		if h[o]++; x < h[o] {
			a, x = []int{o}, h[o]
		} else if x == h[o] {
			a = append(a, o)
		}
		return o
	}
	dfs(root)
	return a
}
