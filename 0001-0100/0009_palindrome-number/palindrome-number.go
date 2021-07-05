package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var flip int
	for i := x; i > 0; i /= 10 {
		flip = flip*10 + i%10
	}
	return flip == x
}
