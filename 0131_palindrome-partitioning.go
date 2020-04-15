package leetcode

/*
  给定一个字符串 s，将 s 分割成一些子串，使每个子串都是回文串。
  返回 s 所有可能的分割方案。
  示例:
    输入: "aab"
    输出:
    [
      ["aa","b"],
      ["a","a","b"]
    ]
*/

// Back Tracking Method
func partition0131(s string) [][]string {
	isPalindrome := func(s string) bool {
		for i := 0; i < len(s)/2; i++ {
			if s[i] != s[len(s)-1-i] {
				return false
			}
		}
		return true
	}
	var out [][]string
	var cur []string
	var recur func(string)
	recur = func(s string) {
		if s == "" {
			tmp := make([]string, len(cur))
			copy(tmp, cur)
			out = append(out, tmp)
			return
		}
		for i := 1; i <= len(s); i++ {
			if !isPalindrome(s[:i]) {
				continue
			}
			cur = append(cur, s[:i])
			recur(s[i:])
			cur = cur[:len(cur)-1]
		}
	}
	recur(s)
	return out
}

// Back Tracking Method && Dynamic Programming
func partition01312(s string) [][]string {
	isPalindrome := func(s string) bool {
		for i := 0; i < len(s)/2; i++ {
			if s[i] != s[len(s)-1-i] {
				return false
			}
		}
		return true
	}
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
	var backTrack func(int)
	backTrack = func(left int) {
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
			backTrack(right)
			cur = cur[:len(cur)-1]
		}
	}
	backTrack(0)
	return out
}

// Back Tracking Method && Dynamic Programming
func partition01313(s string) [][]string {
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
	var backTrack func(int)
	backTrack = func(left int) {
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
			backTrack(right + 1)
			cur = cur[:len(cur)-1]
		}
	}
	backTrack(0)
	return out
}
