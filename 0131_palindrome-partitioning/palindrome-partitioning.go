package main

/*
  给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。
  返回 s 所有可能的分割方案。
*/

// Backtracking
func partition(s string) [][]string {
	var out [][]string
	var cur []string
	var backtrack func(int)
	backtrack = func(offset int) {
		if offset == len(s) {
			tmp := make([]string, len(cur))
			copy(tmp, cur)
			out = append(out, tmp)
			return
		}
		for i := offset; i < len(s); i++ {
			if isPalindrome(s[offset : i+1]) {
				cur = append(cur, s[offset:i+1])
				backtrack(i + 1)
				cur = cur[:len(cur)-1]
			}
		}
	}
	backtrack(0)
	return out
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)>>1; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
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
