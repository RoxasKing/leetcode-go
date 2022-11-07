package main

// Difficulty:
// Medium

// Tags:
// Two Pointers

func magicalString(n int) int {
	o, a := 0, make([]int, n)
	for i, j, x, c := 0, 0, 1, 1; i < n; i++ {
		a[i] = x
		if x == 1 {
			o++
		}
		if c < a[j] {
			c++
			continue
		}
		j++
		x ^= 0b11
		c = 1
	}
	return o
}
