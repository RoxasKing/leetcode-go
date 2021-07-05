package main

// Tags:
// Bit Manipulation
func singleNumbers(nums []int) []int {
	xor := 0
	for _, num := range nums {
		xor ^= num
	}
	flg := 1
	for xor&flg == 0 {
		flg <<= 1
	}
	a, b := 0, 0
	for _, num := range nums {
		if num&flg == flg {
			a ^= num
		} else {
			b ^= num
		}
	}
	return []int{a, b}
}
