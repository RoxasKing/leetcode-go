package main

func sumOfUnique(nums []int) int {
	out := 0
	cnt := [101]int{}
	for _, num := range nums {
		if cnt[num] == 0 {
			out += num
		} else if cnt[num] == 1 {
			out -= num
		}
		cnt[num]++
	}
	return out
}
