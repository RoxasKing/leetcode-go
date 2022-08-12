package main

import (
	"strconv"
)

// Difficulty:
// Medium

// Tags:
// String

func solveEquation(equation string) string {
	isNumber := func(ch byte) bool { return '0' <= ch && ch <= '9' }
	i, n := 0, len(equation)
	parseInt := func() int {
		o := 0
		for ; i < n && isNumber(equation[i]); i++ {
			o = o*10 + int(equation[i]-'0')
		}
		return o
	}
	h := map[byte]int{'+': 1, '-': -1}
	x, v := 0, 0
	for sign := 1; equation[i] != '='; {
		if equation[i] == '+' || equation[i] == '-' {
			sign = h[equation[i]]
			i++
			continue
		}
		num := 1
		if isNumber(equation[i]) {
			num = parseInt()
		}
		if equation[i] == 'x' {
			x += sign * num
			i++
		} else {
			v += sign * num
		}
	}
	i++
	for sign := 1; i < n; {
		if equation[i] == '+' || equation[i] == '-' {
			sign = h[equation[i]]
			i++
			continue
		}
		num := 1
		if isNumber(equation[i]) {
			num = parseInt()
		}
		if i < n && equation[i] == 'x' {
			x -= sign * num
			i++
		} else {
			v -= sign * num
		}
	}
	if x == 0 && v != 0 {
		return "No solution"
	}
	if x == 0 {
		return "Infinite solutions"
	}
	return "x=" + strconv.Itoa(-v/x)
}
