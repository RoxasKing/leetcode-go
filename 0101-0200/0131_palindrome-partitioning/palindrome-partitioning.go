package main

/*
  Given a string s, partition s such that every substring of the partition is a palindrome. Return all possible palindrome partitioning of s.

  A palindrome string is a string that reads the same backward as forward.

  Example 1:
    Input: s = "aab"
    Output: [["a","a","b"],["aa","b"]]

  Example 2:
    Input: s = "a"
    Output: [["a"]]

  Constraints:
    1. 1 <= s.length <= 16
    2. s contains only lowercase English letters.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/palindrome-partitioning
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

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
