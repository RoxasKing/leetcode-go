package main

/*
  删除最小数量的无效括号，使得输入的字符串有效，返回所有可能的结果。

  说明: 输入可能包含了除 ( 和 ) 以外的字符。
*/

// Backtracking + Hash
func removeInvalidParentheses(s string) []string {
	var out []string
	var q []byte
	var lCount, rCount int
	mark := make(map[string]bool)
	var backtrack func(int)
	backtrack = func(index int) {
		if lCount == rCount {
			str := string(q)
			if len(out) == 0 || len(out[0]) < len(str) {
				mark = make(map[string]bool)
				mark[str] = true
				out = []string{str}
			} else if len(out[0]) == len(str) && !mark[str] {
				out = append(out, str)
				mark[str] = true
			}
		}
		if index == len(s) {
			return
		}

		backtrack(index + 1)

		e := s[index]
		if e == '(' {
			lCount++
		} else if e == ')' {
			if lCount == rCount {
				return
			}
			rCount++
		}

		q = append(q, e)
		backtrack(index + 1)
		q = q[:len(q)-1]

		if e == '(' {
			lCount--
		} else if e == ')' {
			rCount--
		}
	}
	backtrack(0)
	return out
}

// Backtracking
func removeInvalidParentheses2(s string) []string {
	var out []string
	var backtrack func(string, int, int, int)
	backtrack = func(s string, start, l, r int) {
		if l == 0 && r == 0 {
			if isValid(s) {
				out = append(out, s)
			}
			return
		}
		for i := start; i < len(s); i++ {
			if s[i] == '(' || s[i] == ')' {
				if i > start && s[i-1] == s[i] {
					continue
				}
				if r > 0 && s[i] == ')' {
					backtrack(s[:i]+s[i+1:], i, l, r-1)
				}
				if r == 0 && l > 0 && s[i] == '(' {
					backtrack(s[:i]+s[i+1:], i, l-1, r)
				}
			}
		}
	}
	l, r := minDelete(s)
	backtrack(s, 0, l, r)
	return out
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
			r++
			l = 0
		}
	}
	return
}
