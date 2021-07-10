package main

// Tags:
// Backtracking

func constructDistancedSequence(n int) []int {
	used := make([]bool, n+1)
	out := make([]int, (n-1)*2+1)
	dfs(out, used, n, (n-1)*2+1, 0)
	return out
}

func dfs(out []int, used []bool, n, size int, i int) bool {
	if i == size {
		for i := range out {
			if out[i] == 0 {
				return false
			}
		}
		return true
	}

	if out[i] != 0 {
		return dfs(out, used, n, size, i+1)
	}

	for x := n; x > 0; x-- {
		if used[x] || x > 1 && (i+x > size-1 || out[i+x] != 0) {
			continue
		}
		used[x] = true
		out[i] = x
		if x > 1 {
			out[i+x] = x
		}
		if dfs(out, used, n, size, i+1) {
			return true
		}
		out[i] = 0
		if x > 1 {
			out[i+x] = 0
		}
		used[x] = false
	}
	return false
}
