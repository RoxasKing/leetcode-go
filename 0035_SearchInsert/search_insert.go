package main

import "fmt"

func searchInsert(nums []int, target int) int {
	size := len(nums)
	start := 0
	mid := size / 2
	end := size - 1
	for {
		if target > nums[end] {
			return end + 1
		} else if target <= nums[start] {
			return start
		} else {
			if target > nums[mid] {
				start = mid
				mid = (mid + end) / 2
			} else if target < nums[mid] {
				end = mid
				mid = (start + mid) / 2
			} else {
				return mid
			}
		}
		if start == mid {
			return start + 1
		}
		if mid == end {
			return mid + 1
		}
	}
}

func main() {
	nums := []int{1, 3, 5, 6, 8, 9, 10, 11}
	target := 6
	fmt.Println(searchInsert(nums, target))
}
