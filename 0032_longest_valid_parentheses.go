package My_LeetCode_In_Go

/*
  给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。
*/

// Dynamic programming
func longestValidParentheses(s string) int {
	var max int
	dp := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1] >= 2 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
			if dp[i] > max {
				max = dp[i]
			}
		}
	}
	return max
}

// Stack
func longestValidParentheses2(s string) int {
	var max int
	stack := make([]int, 0, len(s)/2)
	stack = append(stack, -1)
	for i, c := range s {
		if c == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else if i-stack[len(stack)-1] > max {
				max = i - stack[len(stack)-1]
			}
		}
	}
	return max
}

// Scan two times
func longestValidParentheses3(s string) int {
	var max, l, r int
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			l++
		} else {
			r++
		}
		if l == r {
			if l<<1 > max {
				max = l << 1
			}
		} else if l < r {
			l, r = 0, 0
		}
	}
	l, r = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			l++
		} else {
			r++
		}
		if l == r {
			if l<<1 > max {
				max = l << 1
			}
		} else if l > r {
			l, r = 0, 0
		}
	}
	return max
}
