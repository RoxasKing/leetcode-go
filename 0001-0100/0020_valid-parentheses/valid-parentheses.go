package main

// Tags:
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
