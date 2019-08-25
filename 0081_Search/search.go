package main

import "fmt"

func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	start := 0
	mid := len(nums) / 2
	end := len(nums) - 1
	for {
		if target == nums[start] || target == nums[mid] || target == nums[end] {
			return true
		}
		if start == mid {
			break
		}
		if nums[start] > nums[mid] {
			// 如果右边有序
			if target < nums[mid] || target > nums[end] {
				// 如果 target 小于中间，或者大于尾部，只可能在左侧
				end = mid
				mid = (start + end) / 2
			} else {
				// target 在右侧
				start = mid
				mid = (start + end) / 2
			}
		} else if nums[start] < nums[mid] {
			// 如果左边有序
			if target > nums[mid] || target < nums[start] {
				// 如果 target 大于中间，或者小于首部，只能在右侧
				start = mid
				mid = (start + end) / 2
			} else {
				// target 在左侧
				end = mid
				mid = (start + end) / 2
			}
		} else {
			// 如果头和中间相等，无法判断旋转点在哪边,去掉头和尾节点
			start++
			end--
		}
	}
	return false
}

func main() {
	nums := []int{2, 5, 6, 0, 0, 1, 2}
	//nums = []int{1, 1, 1, 3, 1}
	nums = []int{1, 3, 1, 1}
	target := 3
	fmt.Println(search(nums, target))
}
