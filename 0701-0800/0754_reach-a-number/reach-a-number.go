package main

// Difficulty:
// Medium

// Tags:
// Math

func reachNumber(target int) int {
	if target < 0 {
		target = -target
	}
	k := 0
	for ; target > 0; target -= k {
		k++
	}
	if target%2 == 0 {
		return k
	}
	return k + 1 + k%2
}
