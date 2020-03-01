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

// Back Tracking Method && Dynamic programming
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
	var backTrack func(int, int)
	backTrack = func(left, right int) {
		if left > right {
			tmp := make([]string, len(cur))
			copy(tmp, cur)
			out = append(out, tmp)
			return
		}
		for i := left; i <= right; i++ {
			if !dp[left][i+1] {
				continue
			}
			cur = append(cur, s[left:i+1])
			backTrack(i+1, right)
			cur = cur[:len(cur)-1]
		}
	}
	backTrack(0, len(s)-1)
	return out
}
