package main

// Difficulty:
// Medium

// Tags:
// DFS
// Memoization

func integerReplacement(n int) int {
	memo = map[int]int{}
	return dfs(n)
}

var memo map[int]int

func dfs(n int) int {
	if n == 1 {
		return 0
	}
	if v, ok := memo[n]; ok {
		return v
	}
	if n&1 == 0 {
		memo[n] = 1 + dfs(n>>1)
	} else {
		memo[n] = 1 + Min(dfs(n-1), dfs(n+1))
	}
	return memo[n]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
