package main

func findMagicIndex(nums []int) int {
	var index int
	for index < len(nums) {
		if nums[index] == index {
			return nums[index]
		} else if nums[index] > index {
			index = nums[index]
		} else {
			index++
		}
	}
	return -1
}
