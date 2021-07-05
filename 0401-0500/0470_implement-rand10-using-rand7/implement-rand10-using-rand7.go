package main

import "math/rand"

// Math + Bit Manipulation
func rand10() int {
	a, b := 7, 7

	for a == 7 { // [1, 6]
		a = rand7()
	}
	for b > 5 { // [1, 5]
		b = rand7()
	}

	if a&1 == 1 { // 50% [0, 1] * 20% [1, 5]
		return b
	}
	return b + 5
}

func rand7() int {
	return rand.Intn(7) + 1
}
