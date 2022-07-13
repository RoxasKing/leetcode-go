package main

// Difficulty:
// Easy

// Tags:
// Dynamic Programming

func fib(n int) int {
	a, b := 0, 1
	for ; n > 0; n-- {
		a, b = b, a+b
	}
	return a
}
