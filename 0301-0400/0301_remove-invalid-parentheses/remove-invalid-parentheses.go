package main

// Tags:
// Backtracking
func removeInvalidParentheses(s string) []string {
	var out []string
	l, r := minDelete(s)
	backtrack(&out, s, 0, l, r)
	return out
}

func backtrack(out *[]string, s string, start, l, r int) {
	if l == 0 && r == 0 {
		if isValid(s) {
			*out = append(*out, s)
		}
		return
	}
	for i := start; i < len(s); i++ {
		if !isParentheses(s[i]) || i > start && s[i-1] == s[i] {
			continue
		}
		if r > 0 && s[i] == ')' {
			backtrack(out, s[:i]+s[i+1:], i, l, r-1)
		}
		if r == 0 && l > 0 && s[i] == '(' {
			backtrack(out, s[:i]+s[i+1:], i, l-1, r)
		}
	}
}

func isParentheses(b byte) bool {
	return b == '(' || b == ')'
}

func isValid(s string) bool {
	var count int
	for i := range s {
		if s[i] == '(' {
			count++
		} else if s[i] == ')' {
			count--
		}
		if count < 0 {
			return false
		}
	}
	return count == 0
}

func minDelete(s string) (l, r int) {
	for i := range s {
		if s[i] == '(' {
			l++
		} else if s[i] == ')' {
			l--
		}
		if l < 0 {
			l = 0
			r++
		}
	}
	return
}
