package main

// Difficulty:
// Medium

// Tags:
// Greedy

func findMinFibonacciNumbers(k int) int {
	f := []int{1, 1}
	for f[len(f)-2]+f[len(f)-1] <= k {
		f = append(f, f[len(f)-2]+f[len(f)-1])
	}
	out := 0
	for i := len(f) - 1; i >= 0 && k != 0; i-- {
		if f[i] <= k {
			k -= f[i]
			out++
		}
	}
	return out
}
