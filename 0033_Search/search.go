package main

import "fmt"

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	start := 0
	mid := len(nums) / 2
	end := len(nums) - 1
	for {
		if target == nums[start] {
			return start
		} else if target == nums[mid] {
			return mid
		} else if target == nums[end] {
			return end
		}
		if start == mid {
			break
		}
		if nums[start] > nums[mid] {
			// 如果右边有序
			if target < nums[mid] || target > nums[end] {
				// 如果 target 小于中间，或者大于尾部,只可能在左侧
				end = mid
				mid = (start + end) / 2
			} else {
				// target 在右侧
				start = mid
				mid = (start + end) / 2
			}
		} else {
			// 如果左边有序
			if target > nums[mid] || target < nums[start] {
				// 如果 target 大于中间，或者小于首部，只可能在右侧
				start = mid
				mid = (start + end) / 2
			} else {
				// target 在左侧
				end = mid
				mid = (start + end) / 2
			}
		}
	}
	return -1
}

func main() {
	//nums := []int{4, 5, 6, 7, 0, 1, 2}
	nums := []int{4, 5, 6, 7, 8, 9, 0, 1, 2}
	target := 0
	fmt.Println(search(nums, target))
}
