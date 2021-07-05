package main

// Tags:
// Stack
func calculate(s string) int {
	out := 0
	ops := []int{1}
	sig := 1
	n := len(s)
	for i := 0; i < n; {
		switch s[i] {
		case ' ':
			i++
		case '+':
			sig = ops[len(ops)-1]
			i++
		case '-':
			sig = -ops[len(ops)-1]
			i++
		case '(':
			ops = append(ops, sig)
			i++
		case ')':
			ops = ops[:len(ops)-1]
			i++
		default:
			num := 0
			for ; i < n && '0' <= s[i] && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			out += sig * num
		}
	}
	return out
}
