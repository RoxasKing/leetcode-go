package main

// Difficulty:
// Medium

// Tags:
// Bit Manipulation
// Backtracking

func grayCode(n int) []int {
	used := make([]bool, 1<<n)
	used[0] = true
	out := []int{0}
	cur := 0
	ok := false
	var dfs func(int)
	dfs = func(k int) {
		if k == (1<<n)-1 {
			ok = true
			return
		}
		for i := 0; i < n; i++ {
			if used[cur^(1<<i)] {
				continue
			}
			cur ^= 1 << i
			used[cur] = true
			out = append(out, cur)
			dfs(k + 1)
			if ok {
				return
			}
			out = out[:len(out)-1]
			used[cur] = false
			cur ^= 1 << i
		}
	}
	dfs(0)
	return out
}
