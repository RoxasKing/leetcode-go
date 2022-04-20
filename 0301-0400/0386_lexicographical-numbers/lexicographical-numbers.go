package main

// Difficulty:
// Medium

// Tags:
// Math

func lexicalOrder(n int) []int {
	o := make([]int, n)
	for i, x := 0, 1; i < n; i++ {
		o[i] = x
		if x*10 <= n {
			x *= 10
			continue
		}
		for x%10 == 9 || x+1 > n {
			x /= 10
		}
		x++
	}
	return o
}
