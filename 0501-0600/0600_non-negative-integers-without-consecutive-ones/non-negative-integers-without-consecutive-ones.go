package main

// Tags:
// Dynamic Programming
// Bit Manipulation

func findIntegers(n int) int {
	f := [31]int{1, 1}
	for i := 2; i < 31; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	out, i := 0, 30
	for ; i >= 1; i-- {
		x := n >> (i - 1)
		if x&1 == 1 {
			out += f[i]
			if x&2 > 0 {
				break
			}
		}
	}
	if i == 0 {
		out++
	}
	return out
}
