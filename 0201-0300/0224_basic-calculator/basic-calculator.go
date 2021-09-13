package main

// Tags:
// Stack

func calculate(s string) int {
	out := 0
	ops := []int{1}
	sig := 1
	n := len(s)
	for i := 0; i < n; {
		ch := s[i]
		if isNumber(ch) {
			num := 0
			for ; i < n && isNumber(s[i]); i++ {
				num = num*10 + int(s[i]-'0')
			}
			out += sig * num
			continue
		}
		if ch == '+' {
			sig = ops[len(ops)-1]
		} else if ch == '-' {
			sig = -ops[len(ops)-1]
		} else if ch == '(' {
			ops = append(ops, sig)
		} else if ch == ')' {
			ops = ops[:len(ops)-1]
		}
		i++
	}
	return out
}

func isNumber(ch byte) bool { return '0' <= ch && ch <= '9' }
