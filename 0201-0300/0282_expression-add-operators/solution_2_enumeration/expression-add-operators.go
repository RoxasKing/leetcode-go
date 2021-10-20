package main

import "math"

// Difficulty:
// Hard

// Tags:
// Enumeration

func addOperators(num string, target int) []string {
	n := len(num)
	ops := make([]byte, n) // 0: ' '; 1: '+'; 2: '-'; 3: '*'
	N := int(math.Pow(float64(4), float64(n-1)))
	out := []string{}
	for i := 0; i < N; i++ {
		for j, x := n-1, i; j > 0; j-- {
			ops[j] = Op(x % 4)
			x /= 4
		}
		ok := true
		stk := []int{}
		chs := make([]byte, 0, n<<1)
		chs = append(chs, num[0])
		t, p := int(num[0]-'0'), byte('+')
		for j := 1; j < n; j++ {
			x := int(num[j] - '0')
			if ops[j] == ' ' {
				if t == 0 {
					ok = false
					break
				}
				t = t*10 + x
				chs = append(chs, num[j])
			} else {
				if p == '+' {
					stk = append(stk, t)
				} else if p == '-' {
					stk = append(stk, -t)
				} else if p == '*' {
					stk[len(stk)-1] *= t
				}
				t = x
				p = ops[j]
				chs = append(chs, ops[j], num[j])
			}
		}
		if !ok {
			continue
		}
		if p == '+' {
			stk = append(stk, t)
		} else if p == '-' {
			stk = append(stk, -t)
		} else if p == '*' {
			stk[len(stk)-1] *= t
		}
		sum := 0
		for _, x := range stk {
			sum += x
		}
		if sum == target {
			out = append(out, string(chs))
		}
	}
	return out
}

func Op(x int) byte {
	if x == 0 {
		return ' '
	} else if x == 1 {
		return '+'
	} else if x == 2 {
		return '-'
	}
	return '*'
}
