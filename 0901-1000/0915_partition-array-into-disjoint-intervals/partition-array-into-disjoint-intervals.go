package main

// Difficulty:
// Medium

func partitionDisjoint(nums []int) int {
	o := 1
	cur, max := nums[0], nums[0]
	for i, x := range nums {
		if x < cur {
			o = i + 1
			cur = max
		} else if max < x {
			max = x
		}
	}
	return o
}
