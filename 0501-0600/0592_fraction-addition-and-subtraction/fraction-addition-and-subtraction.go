package main

import (
	"strconv"
	"strings"
)

// Difficulty:
// Medium

func fractionAddition(expression string) string {
	x, y, sign := 0, 1, 1
	n := len(expression)
	for i := 0; i < n; i++ {
		if expression[i] == '+' {
			sign = 1
			continue
		}
		if expression[i] == '-' {
			sign = -1
			continue
		}
		j := i
		for ; j < n && expression[j] != '+' && expression[j] != '-'; j++ {
		}
		s := strings.Split(expression[i:j], "/")
		i = j - 1
		a, _ := strconv.Atoi(s[0])
		b, _ := strconv.Atoi(s[1])
		if x == 0 {
			x, y = sign*a, b
			continue
		}
		g := gcd(b, y)
		l := y * b / g
		x, y = x*l/y+sign*a*l/b, l
		if g := abs(gcd(x, y)); g > 1 {
			x, y = x/g, y/g
		}
	}
	return strconv.Itoa(x) + "/" + strconv.Itoa(y)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
