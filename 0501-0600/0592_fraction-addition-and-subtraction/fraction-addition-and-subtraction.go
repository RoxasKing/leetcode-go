package main

import (
	"strconv"
	"strings"
)

// Difficulty:
// Medium

// Tags:
// Math
// Simulation

func fractionAddition(expression string) string {
	numerator, denominator := 0, 1
	h := map[byte]int{'+': 1, '-': -1}
	for i, n, sign := 0, len(expression), 1; i < n; {
		if expression[i] == '+' || expression[i] == '-' {
			sign = h[expression[i]]
			i++
			continue
		}
		i0 := i
		for ; i < n && expression[i] != '+' && expression[i] != '-'; i++ {
		}
		a := strings.Split(expression[i0:i], "/")
		x, _ := strconv.Atoi(a[0])
		y, _ := strconv.Atoi(a[1])
		lcm := abs(denominator*y) / gcd(denominator, y)
		numerator, denominator = numerator*lcm/denominator+sign*x*lcm/y, lcm
		if g := abs(gcd(numerator, denominator)); g > 1 {
			numerator, denominator = numerator/g, denominator/g
		}
	}
	return strconv.Itoa(numerator) + "/" + strconv.Itoa(denominator)
}

func gcd(a, b int) int {
	for ; b != 0; a, b = b, a%b {
	}
	return a
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
