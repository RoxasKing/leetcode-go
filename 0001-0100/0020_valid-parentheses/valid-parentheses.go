package main

/*
  Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

  An input string is valid if:
    1. Open brackets must be closed by the same type of brackets.
    2. Open brackets must be closed in the correct order.

  Example 1:
    Input: s = "()"
    Output: true

  Example 2:
    Input: s = "()[]{}"
    Output: true

  Example 3:
    Input: s = "(]"
    Output: false

  Example 4:
    Input: s = "([)]"
    Output: false

  Example 5:
    Input: s = "{[]}"
    Output: true

  Constraints:
    1. 1 <= s.length <= 10^4
    2. s consists of parentheses only '()[]{}'.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/valid-parentheses
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Stack
func isValid(s string) bool {
	stk := []byte{}
	for i := range s {
		if s[i] == '[' || s[i] == '(' || s[i] == '{' {
			stk = append(stk, s[i])
			continue
		}
		if len(stk) == 0 {
			return false
		}
		top := len(stk) - 1
		ch := stk[top]
		if ch == '[' && s[i] != ']' || ch == '(' && s[i] != ')' || ch == '{' && s[i] != '}' {
			return false
		}
		stk = stk[:top]
	}
	return len(stk) == 0
}
