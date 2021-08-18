package main

import "math/bits"

// Tags:
// Dynamic Programming

func countArrangement(n int) int {
	f := [1 << 15]int{1}
	for mask := 1; mask < 1<<n; mask++ {
		i := bits.OnesCount(uint(mask))
		for j := 0; j < n; j++ {
			if (mask>>j)&1 == 1 && (i%(j+1) == 0 || (j+1)%i == 0) {
				f[mask] += f[mask^(1<<j)]
			}
		}
	}
	return f[1<<n-1]
}
