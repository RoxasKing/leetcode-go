package main

// Tags:
// Bit Manipulation
func singleNumber(nums []int) int {
	out := 0
	for i := 1; i <= 1<<31; i <<= 1 {
		cnt := 0
		for _, num := range nums {
			if num&i == i {
				cnt++
			}
		}
		if cnt%3 == 1 {
			out |= i
		}
	}
	return out
}
