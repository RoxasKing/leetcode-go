package main

func findRepeatNumber(nums []int) int {
	for i, num := range nums {
		j := num
		for i != j {
			if nums[i] == nums[j] {
				return j
			}
			nums[i], nums[j] = nums[j], nums[i]
			j = nums[i]
		}
	}
	return -1
}
