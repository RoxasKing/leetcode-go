package main

import "math"

// Math
// Lagrange's four-square theorem

func numSquares(n int) int {
	for n%4 == 0 {
		n /= 4
	}
	if n%8 == 7 {
		return 4
	}
	for a := 0; a*a <= n; a++ {
		b := int(math.Pow(float64(n-a*a), 0.5))
		if a*a+b*b == n {
			if a != 0 && b != 0 {
				return 2
			}
			return 1
		}
	}
	return 3
}
