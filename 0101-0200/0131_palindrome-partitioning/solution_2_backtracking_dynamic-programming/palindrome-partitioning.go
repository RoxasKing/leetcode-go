package main

// Difficulty:
// Medium

// Tags:
// Backtracking
// Dynamic Programming

func partition(s string) [][]string {
	isPalindrome := func(l, r int) bool {
		for l < r {
			if s[l] != s[r] {
				return false
			}
			l++
			r--
		}
		return true
	}
	n := len(s)
	f := make([][]bool, n)
	for i := range f {
		f[i] = make([]bool, n)
		f[i][i] = true
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isPalindrome(i, j) {
				f[i][j] = true
			}
		}
	}
	var out [][]string
	var cur []string
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			tmp := make([]string, len(cur))
			copy(tmp, cur)
			out = append(out, tmp)
			return
		}
		for j := i; j < n; j++ {
			if !f[i][j] {
				continue
			}
			cur = append(cur, s[i:j+1])
			dfs(j + 1)
			cur = cur[:len(cur)-1]
		}
	}
	dfs(0)
	return out
}
