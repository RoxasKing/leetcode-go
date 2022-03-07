package main

// Difficulty:
// Hard

// Tags:
// Backtracking

func removeInvalidParentheses(s string) []string {
	n := len(s)
	del := make([]bool, n)
	out := []string{}
	l, r := 0, 0
	for i := range s {
		if c := s[i]; c == '(' {
			l++
		} else if c == ')' {
			l--
		}
		if l < 0 {
			l = 0
			r++
		}
	}
	tmp := make([]byte, n-l-r)
	var backtrack func(int, int, int)
	backtrack = func(idx, l, r int) {
		if l == 0 && r == 0 {
			cnt, idx := 0, 0
			for i := range s {
				if del[i] {
					continue
				}
				c := s[i]
				if c == '(' {
					cnt++
				} else if c == ')' {
					if cnt--; cnt < 0 {
						break
					}
				}
				tmp[idx] = c
				idx++
			}
			if cnt == 0 {
				out = append(out, string(tmp))
			}
			return
		}
		for i := idx; i < n; i++ {
			if s[i] != '(' && s[i] != ')' || i > idx && s[i-1] == s[i] {
				continue
			}
			if s[i] == '(' && r == 0 && l > 0 {
				del[i] = true
				backtrack(i+1, l-1, r)
				del[i] = false
			} else if s[i] == ')' && r > 0 {
				del[i] = true
				backtrack(i+1, l, r-1)
				del[i] = false
			}
		}
	}
	backtrack(0, l, r)
	return out
}
