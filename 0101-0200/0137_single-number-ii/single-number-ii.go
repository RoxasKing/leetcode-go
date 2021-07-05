package main

// Tags:
// Bit Manipulation
func singleNumber(nums []int) int {
	var s1, s2 int
	for _, num := range nums {
		s1 = ^s2 & (s1 ^ num)
		s2 = ^s1 & (s2 ^ num)
	}
	return s1
}
