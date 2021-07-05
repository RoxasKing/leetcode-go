package main

// Tags:
// Math
func checkPowersOfThree(n int) bool {
	for n > 0 {
		remain := n % 3
		if remain != 0 && remain != 1 {
			return false
		}
		n /= 3
	}
	return true
}
