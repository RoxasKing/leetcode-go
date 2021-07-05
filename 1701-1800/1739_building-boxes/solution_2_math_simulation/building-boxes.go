package main

// Tags:
// Math + Simulation

func minimumBoxes(n int) int {
	if n <= 3 {
		return n
	}
	out := 0
	i, upper, sum := 1, 1, 0
	for sum < n {
		out++
		sum += i
		i++
		if i > upper {
			i = 1
			upper++
		}
	}
	return out
}
