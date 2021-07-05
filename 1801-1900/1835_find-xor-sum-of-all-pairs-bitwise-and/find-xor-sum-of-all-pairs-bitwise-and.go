package main

// Tags:
// Bit Manipulation
func getXORSum(arr1 []int, arr2 []int) int {
	xor1 := 0
	for _, num := range arr1 {
		xor1 ^= num
	}
	xor2 := 0
	for _, num := range arr2 {
		xor2 ^= num
	}
	return xor1 & xor2
}
