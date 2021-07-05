package main

// Tags:
// Dynamic Programming

func nthSuperUglyNumber(n int, primes []int) int {
	pointers := make([]int, len(primes))
	for i := 1; i < n; i++ {
		min := 1<<31 - 1
		for j, p := range pointers {
			min = Min(min, f[p]*primes[j])
		}
		f[i] = min
		for j, p := range pointers {
			if f[p]*primes[j] <= min {
				pointers[j]++
			}
		}
	}
	return f[n-1]
}

var f = [100000]int{1}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
