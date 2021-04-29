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

// Stack
func longestValidParentheses(s string) int {
	out := 0
	stk := []int{-1}
	for i := range s {
		if s[i] == '(' {
			stk = append(stk, i)
			continue
		}

		if len(stk) == 1 {
			stk[0] = i
			continue
		}

		stk = stk[:len(stk)-1]
		out = Max(out, i-stk[len(stk)-1])
	}

	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
