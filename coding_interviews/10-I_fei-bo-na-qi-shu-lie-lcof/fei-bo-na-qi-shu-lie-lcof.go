package main

// Tags:
// Dynamic Programming
func fib(n int) int {
	a, b := 0, 1
	for i := 1; i <= n; i++ {
		a, b = b, (a+b)%(1e9+7)
	}
	return a
}
