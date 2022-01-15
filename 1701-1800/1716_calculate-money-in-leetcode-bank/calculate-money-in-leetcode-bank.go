package main

// Difficulty:
// Easy

// Tags:
// Math

func totalMoney(n int) int {
	out := 0
	a, b := n/7, n%7
	out += 28*a + a*(a-1)*7/2
	out += (1+a)*b + b*(b-1)/2
	return out
}
