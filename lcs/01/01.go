package main

// Tags:
// Greedy Algorithm

func leastMinutes(n int) int {
	out := 1
	i := 1
	for i < n {
		out++
		i <<= 1
	}
	return out
}
