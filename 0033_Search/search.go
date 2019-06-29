package main

import "fmt"

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	start := 0
	mid := len(nums) / 2
	end := len(nums) - 1
	// 寻找旋转点
	for {
		if nums[mid] >= nums[start] && nums[mid] >= nums[end] {

		}
	}
	return -1
}

func main() {
	//nums := []int{4, 5, 6, 7, 0, 1, 2}
	nums := []int{4, 5, 6, 7, 8, 9, 0, 1, 2}
	target := 9
	fmt.Println(search(nums, target))
}
