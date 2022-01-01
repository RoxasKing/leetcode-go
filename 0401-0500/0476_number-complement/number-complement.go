package main

// Difficulty:
// Easy

// Tags:
// Bit Manipulation

func findComplement(num int) int {
	mask := 0
	for x := num; x != 0; x >>= 1 {
		mask = mask<<1 + 1
	}
	return mask ^ num
}
