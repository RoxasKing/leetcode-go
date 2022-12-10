package main

// Difficulty:
// Medium

// Tags:
// Math

func checkPowersOfThree(n int) bool {
	for ; n > 0; n /= 3 {
		if rem := n % 3; rem != 0 && rem != 1 {
			return false
		}
	}
	return true
}
