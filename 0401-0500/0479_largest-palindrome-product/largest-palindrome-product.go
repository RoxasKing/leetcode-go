package main

import "math"

// Difficulty:
// Hard

// Tags:
// Math
// Enumeration

func largestPalindrome(n int) int {
	if n == 1 {
		return 9
	}
	max := int(math.Pow10(n)) - 1
	for i := max; ; i-- {
		num := i
		for x := i; x > 0; x /= 10 {
			num = num*10 + x%10
		}
		for x := max; x*x >= num; x-- {
			if num%x == 0 {
				return num % 1337
			}
		}
	}
}
