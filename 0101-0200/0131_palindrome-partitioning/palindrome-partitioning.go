package main

// Tags:
// Backtracking
func partition(s string) [][]string {
	out := [][]string{}
	dfs(s, []string{}, &out)
	return out
}

func dfs(s string, cur []string, out *[][]string) {
	if s == "" {
		tmp := make([]string, len(cur))
		copy(tmp, cur)
		*out = append(*out, tmp)
		return
	}

	for i := range s {
		if isPalindrome(s[:i+1]) {
			cur = append(cur, s[:i+1])
			dfs(s[i+1:], cur, out)
			cur = cur[:len(cur)-1]
		}
	}
}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

// Backtracking && Dynamic Programming
func partition2(s string) [][]string {
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s)+1)
		for j := i + 1; j <= len(s); j++ {
			if isPalindrome(s[i:j]) {
				dp[i][j] = true
			}
		}
	}
	var out [][]string
	var cur []string
	var backtrack func(int)
	backtrack = func(left int) {
		if left == len(s) {
			tmp := make([]string, len(cur))
			copy(tmp, cur)
			out = append(out, tmp)
			return
		}
		for right := left + 1; right <= len(s); right++ {
			if !dp[left][right] {
				continue
			}
			cur = append(cur, s[left:right])
			backtrack(right)
			cur = cur[:len(cur)-1]
		}
	}
	backtrack(0)
	return out
}

// Backtracking && Dynamic Programming
func partition3(s string) [][]string {
	check := make([][]bool, len(s))
	for r := range check {
		check[r] = make([]bool, len(s))
		for l := 0; l <= r; l++ {
			if s[l] == s[r] && (r-l <= 2 || check[l+1][r-1]) {
				check[l][r] = true
			}
		}
	}
	var out [][]string
	var cur []string
	var backtrack func(int)
	backtrack = func(left int) {
		if left == len(s) {
			tmp := make([]string, len(cur))
			copy(tmp, cur)
			out = append(out, tmp)
			return
		}
		for right := left; right < len(s); right++ {
			if !check[left][right] {
				continue
			}
			cur = append(cur, s[left:right+1])
			backtrack(right + 1)
			cur = cur[:len(cur)-1]
		}
	}
	backtrack(0)
	return out
}
