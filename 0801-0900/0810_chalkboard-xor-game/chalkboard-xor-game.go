package main

// Tags:
// Math
// Bit Manipulation

func xorGame(nums []int) bool {
	if len(nums)&1 == 0 {
		return true
	}
	xor := 0
	for _, num := range nums {
		xor ^= num
	}
	return xor == 0
}
