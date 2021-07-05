package main

// Tags:
// Bit Manipulation
func getMaximumXor(nums []int, maximumBit int) []int {
	mask := 1<<maximumBit - 1
	xor := 0
	for _, num := range nums {
		xor ^= num
	}
	n := len(nums)
	out := make([]int, n)
	for i := 0; i < n; i++ {
		out[i] = (xor | mask) ^ xor
		xor ^= nums[n-1-i]
	}
	return out
}
