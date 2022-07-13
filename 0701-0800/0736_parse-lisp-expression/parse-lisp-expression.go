package main

// Difficulty:
// Hard

// Tags:
// Recursion
// Hash
// Stack

func evaluate(expression string) int {
	i, n := 0, len(expression)
	parseStr := func() string {
		bg := i
		for ; i < n && expression[i] != ' ' && expression[i] != ')'; i++ {
		}
		return expression[bg:i]
	}
	parseInt := func() int {
		sign, v := 1, 0
		if expression[i] == '-' {
			sign = -1
			i++
		}
		for ; i < n && '0' <= expression[i] && expression[i] <= '9'; i++ {
			v = v*10 + int(expression[i]-'0')
		}
		return sign * v
	}
	stk := map[string][]int{}
	top := func(k string) int { return stk[k][len(stk[k])-1] }
	var innerEvaluate func() int
	innerEvaluate = func() int {
		var o int
		if expression[i] != '(' {
			if 'a' <= expression[i] && expression[i] <= 'z' {
				return top(parseStr())
			}
			return parseInt()
		}
		i++
		switch expression[i] {
		case 'l':
			i += 4
			a := []string{}
			for {
				if !('a' <= expression[i] && expression[i] <= 'z') {
					o = innerEvaluate()
					break
				}
				e := parseStr()
				if expression[i] == ')' {
					o = top(e)
					break
				}
				a = append(a, e)
				i++
				stk[e] = append(stk[e], innerEvaluate())
				i++
			}
			for _, e := range a {
				stk[e] = stk[e][:len(stk[e])-1]
			}
		case 'a':
			i += 4
			e1 := innerEvaluate()
			i++
			e2 := innerEvaluate()
			o = e1 + e2
		case 'm':
			i += 5
			e1 := innerEvaluate()
			i++
			e2 := innerEvaluate()
			o = e1 * e2
		}
		i++
		return o
	}
	return innerEvaluate()
}
