package main

// Difficulty:
// Easy

// Tags:
// Stack

func isValid(s string) bool {
	stk := []byte{}
	for i := range s {
		c := s[i]
		if c == '(' || c == '[' || c == '{' {
			stk = append(stk, c)
			continue
		}
		if len(stk) == 0 {
			return false
		}
		if t := stk[len(stk)-1]; t == '(' && c != ')' || t == '[' && c != ']' || t == '{' && c != '}' {
			return false
		}
		stk = stk[:len(stk)-1]
	}
	return len(stk) == 0
}
