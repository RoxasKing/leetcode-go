package main

// Tags:
// DFS

func minSteps(n int) int {
	out := 0
	for i := 2; i*i <= n; i++ {
		for ; n%i == 0; n /= i {
			out += i
		}
	}
	if n > 1 {
		out += n
	}
	return out
}
