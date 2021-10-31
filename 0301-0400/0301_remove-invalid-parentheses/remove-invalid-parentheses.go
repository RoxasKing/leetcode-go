package main

// Difficulty:
// Hard

// Tags:
// Backtracking

func removeInvalidParentheses(s string) []string {
	out := []string{}
	l, r := minDelete(s)
	dfs(&out, s, 0, l, r)
	return out
}

func dfs(out *[]string, s string, start, l, r int) {
	if l == 0 && r == 0 {
		if isValid(s) {
			*out = append(*out, s)
		}
		return
	}
	for i := start; i < len(s); i++ {
		if s[i] != '(' && s[i] != ')' || i > start && s[i-1] == s[i] {
			continue
		}
		if r == 0 && l > 0 && s[i] == '(' {
			dfs(out, s[:i]+s[i+1:], i, l-1, r)
		}
		if r > 0 && s[i] == ')' {
			dfs(out, s[:i]+s[i+1:], i, l, r-1)
		}
	}
}

func isValid(s string) bool {
	cnt := 0
	for i := range s {
		if s[i] == '(' {
			cnt++
		} else if s[i] == ')' {
			cnt--
		}
		if cnt < 0 {
			return false
		}
	}
	return cnt == 0
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
