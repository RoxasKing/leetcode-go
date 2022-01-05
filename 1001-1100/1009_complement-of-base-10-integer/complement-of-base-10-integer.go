package main

// Difficulty:
// Easy

// Tags:
// Bit Manipulation

func bitwiseComplement(n int) int {
	if n == 0 {
		return 1
	}
	mask := 0
	for x := n; x > 0; x >>= 1 {
		mask = mask<<1 + 1
	}
	return mask ^ n
}
