package main

// Difficulty:
// Hard

// Tags:
// Stack

func parseBoolExpr(expression string) bool {
	stk := []rune{}
	top := func() int { return len(stk) - 1 }
	for _, c := range expression {
		if c == ',' {
			continue
		}
		if c != ')' {
			stk = append(stk, c)
			continue
		}
		t, f := 0, 0
		for ; stk[top()] != '('; stk = stk[:top()] {
			if v := stk[top()]; v == 't' {
				t++
			} else {
				f++
			}
		}
		stk = stk[:top()]
		op := stk[top()]
		stk = stk[:top()]
		c = 't'
		switch op {
		case '!':
			if f != 1 {
				c = 'f'
			}
			stk = append(stk, c)
		case '&':
			if f != 0 {
				c = 'f'
			}
			stk = append(stk, c)
		case '|':
			if t == 0 {
				c = 'f'
			}
			stk = append(stk, c)
		}
	}
	return stk[top()] == 't'
}
