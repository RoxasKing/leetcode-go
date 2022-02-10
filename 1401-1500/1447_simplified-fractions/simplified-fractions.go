package main

import "strconv"

// Difficulty:
// Medium

// Tags:
// Math

func simplifiedFractions(n int) []string {
	out := []string{}
	for i := 1; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if gcd(i, j) > 1 {
				continue
			}
			out = append(out, strconv.Itoa(i)+"/"+strconv.Itoa(j))
		}
	}
	return out
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
