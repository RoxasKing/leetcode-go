package main

// Difficulty:
// Easy

// Tags:
// Bit Manipulation

func binaryGap(n int) int {
	o := 0
	for c := 0; n > 0; n >>= 1 {
		if n&1 == 1 {
			if o < c {
				o = c
			}
			c = 1
		} else if c > 0 {
			c++
		}
	}
	return o
}
