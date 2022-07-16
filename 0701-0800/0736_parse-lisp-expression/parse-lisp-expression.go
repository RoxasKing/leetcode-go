package main

// Difficulty:
// Hard

// Tags:
// Recursion
// Hash
// Stack

func evaluate(expression string) int {
	isNumber := func(ch byte) bool { return '0' <= ch && ch <= '9' }
	isLetter := func(ch byte) bool { return 'a' <= ch && ch <= 'z' }
	i := 0
	parseStr := func() string {
		i0 := i
		for ; expression[i] != ' ' && expression[i] != ')'; i++ {
		}
		return expression[i0:i]
	}
	parseInt := func() int {
		sign, v := 1, 0
		if expression[i] == '-' {
			sign = -1
			i++
		}
		for ; isNumber(expression[i]); i++ {
			v = v*10 + int(expression[i]-'0')
		}
		return sign * v
	}
	h := map[string][]int{}
	getV := func(k string) int { return h[k][len(h[k])-1] }
	var innerEvaluate func() int
	innerEvaluate = func() int {
		var o int
		if expression[i] != '(' {
			if isLetter(expression[i]) {
				return getV(parseStr())
			}
			return parseInt()
		}
		i++
		switch expression[i] {
		case 'l':
			i += 4
			a := []string{}
			for {
				if !isLetter(expression[i]) {
					o = innerEvaluate()
					break
				}
				k := parseStr()
				if expression[i] == ')' {
					o = getV(k)
					break
				}
				a = append(a, k)
				i++
				h[k] = append(h[k], innerEvaluate())
				i++
			}
			for _, k := range a {
				h[k] = h[k][:len(h[k])-1]
			}
		case 'a':
			i += 4
			v1 := innerEvaluate()
			i++
			v2 := innerEvaluate()
			o = v1 + v2
		case 'm':
			i += 5
			v1 := innerEvaluate()
			i++
			v2 := innerEvaluate()
			o = v1 * v2
		}
		i++
		return o
	}
	return innerEvaluate()
}
