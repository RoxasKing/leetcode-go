package main

// Tags:
// Math

func countTriples(n int) int {
	var out int

	exist := make([]bool, n*n+1)
	for i := 1; i <= n; i++ {
		exist[i*i] = true
	}

	for a := 2; a <= n; a++ {
		for b := 1; b < a; b++ {
			if sum := a*a + b*b; sum <= n*n && exist[sum] {
				out += 2
			}
		}
	}

	return out
}
