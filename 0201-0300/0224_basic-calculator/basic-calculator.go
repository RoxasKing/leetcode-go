package main

// Difficulty:
// Hard

// Tags:
// Stack

func calculate(s string) int {
	isNumber := func(ch byte) bool { return '0' <= ch && ch <= '9' }
	stk := []int{1}
	top := func() int { return len(stk) - 1 }
	o := 0
	for i, n, sign := 0, len(s), 1; i < n; {
		ch := s[i]
		if isNumber(ch) {
			x := 0
			for ; i < n && isNumber(s[i]); i++ {
				x = x*10 + int(s[i]-'0')
			}
			o += sign * x
			continue
		}
		if ch == '+' {
			sign = stk[top()]
		} else if ch == '-' {
			sign = -stk[top()]
		} else if ch == '(' {
			stk = append(stk, sign)
		} else if ch == ')' {
			stk = stk[:top()]
		}
		i++
	}
	return o
}
