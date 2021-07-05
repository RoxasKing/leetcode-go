package main

func removeDuplicates(nums []int) int {
	prev := -1 << 31
	idx := 0
	cnt := 0
	for _, num := range nums {
		if num != prev {
			prev = num
			nums[idx] = num
			idx++
			cnt = 1
		} else if cnt < 2 {
			nums[idx] = num
			idx++
			cnt++
		}
	}
	return idx
}
