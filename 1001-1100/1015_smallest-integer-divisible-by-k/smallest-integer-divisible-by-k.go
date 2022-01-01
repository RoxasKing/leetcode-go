package main

// Difficulty:
// Medium

// Tags:
// Math

func smallestRepunitDivByK(k int) int {
	if k%2 == 0 || k%5 == 0 {
		return -1
	}
	n := 1
	for r := 1; ; r = 10*r + 1 {
		if r %= k; r == 0 {
			break
		}
		n++
	}
	return n
}
