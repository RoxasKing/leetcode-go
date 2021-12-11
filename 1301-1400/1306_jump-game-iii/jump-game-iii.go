package main

// Difficulty:
// Medium

// Tags:
// DFS

func canReach(arr []int, start int) bool {
	vis := make([]bool, len(arr))
	var dfs func(int) bool
	dfs = func(i int) bool {
		if vis[i] {
			return false
		}
		vis[i] = true
		defer func() { vis[i] = false }()
		if arr[i] == 0 {
			return true
		}
		if i+arr[i] < len(arr) && dfs(i+arr[i]) || i-arr[i] >= 0 && dfs(i-arr[i]) {
			return true
		}
		return false
	}
	return dfs(start)
}
