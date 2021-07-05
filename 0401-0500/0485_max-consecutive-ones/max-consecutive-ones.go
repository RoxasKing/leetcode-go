package main

func findMaxConsecutiveOnes(nums []int) int {
	out := 0
	cnt := 0
	for _, num := range nums {
		if num == 1 {
			cnt++
			if cnt > out {
				out = cnt
			}
		} else {
			cnt = 0
		}
	}
	return out
}
