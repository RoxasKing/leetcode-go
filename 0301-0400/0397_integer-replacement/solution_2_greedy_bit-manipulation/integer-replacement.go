package main

// Difficulty:
// Medium

// Tags:
// Greedy
// Bit Manipulation

func integerReplacement(n int) int {
	out := 0
	for ; n > 1; out++ {
		if n&1 == 0 {
			n >>= 1
		} else if n > 0b11 && n&0b11 == 0b11 {
			n++
		} else {
			n--
		}
	}
	return out
}
