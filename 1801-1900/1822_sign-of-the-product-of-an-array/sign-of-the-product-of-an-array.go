package main

func arraySign(nums []int) int {
	out := 1
	for _, num := range nums {
		if num == 0 {
			return 0
		} else if num < 0 {
			out = -out
		}
	}
	return out
}
