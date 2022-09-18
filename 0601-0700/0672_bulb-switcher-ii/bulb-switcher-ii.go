package main

// Difficulty:
// Medium

// Tags:
// Math

func flipLights(n int, presses int) int {
	if n > 3 {
		n = 3
	}
	if presses == 0 {
		return 1
	} else if presses == 1 {
		return []int{2, 3, 4}[n-1]
	} else if presses == 2 {
		return []int{2, 4, 7}[n-1]
	}
	return []int{2, 4, 8}[n-1]
}
