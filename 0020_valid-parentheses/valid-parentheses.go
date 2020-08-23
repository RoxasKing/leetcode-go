package main

/*
  给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
  有效字符串需满足：
    左括号必须用相同类型的右括号闭合。
    左括号必须以正确的顺序闭合。
  注意空字符串可被认为是有效字符串。
*/

func isValid(s string) bool {
	var stack []byte
	for i := range s {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
		} else if len(stack) != 0 &&
			(s[i] == ')' && stack[len(stack)-1] == '(' ||
				s[i] == '}' && stack[len(stack)-1] == '{' ||
				s[i] == ']' && stack[len(stack)-1] == '[') {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}
