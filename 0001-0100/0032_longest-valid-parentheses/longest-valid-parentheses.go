package main

/*
  给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。

  示例 1:
    输入: "(()"
    输出: 2
    解释: 最长有效括号子串为 "()"

  示例 2:
    输入: ")()())"
    输出: 4
    解释: 最长有效括号子串为 "()()"

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/longest-valid-parentheses
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Scan two times
func longestValidParentheses(s string) int {
	var max int
	lCount, rCount := 0, 0
	for i := range s {
		if s[i] == '(' {
			lCount++
		} else if s[i] == ')' {
			rCount++
			if lCount == rCount {
				max = Max(max, lCount+rCount)
			} else if lCount < rCount {
				lCount, rCount = 0, 0
			}
		}
	}
	lCount, rCount = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			lCount++
			if lCount == rCount {
				max = Max(max, lCount+rCount)
			} else if lCount > rCount {
				lCount, rCount = 0, 0
			}
		} else if s[i] == ')' {
			rCount++
		}
	}
	return max
}

// Stack
func longestValidParentheses2(s string) int {
	var max int
	stack := []int{-1}
	for i, c := range s {
		if c == '(' {
			stack = append(stack, i)
			continue
		}
		if len(stack) == 1 {
			stack[0] = i
			continue
		}
		stack = stack[:len(stack)-1]
		max = Max(max, i-stack[len(stack)-1])
	}
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
