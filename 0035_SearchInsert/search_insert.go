package main

import "fmt"

func searchInsert(nums []int, target int) int {
	size := len(nums)
	head := 0
	mid := size / 2
	tail := size - 1
	if target > nums[tail] {
		// 如果目标值大于数组尾
		return tail + 1
	} else if target <= nums[head] {
		// 如果目标值小于数组头
		return head
	}
	for {
		// 如果目标值在数组范围之内
		if target > nums[mid] {
			// 目标值大于中间值,继续检索右半边
			head = mid
			mid = (mid + tail) / 2
		} else if target < nums[mid] {
			// 目标值小于中间值,继续检索左半边
			tail = mid
			mid = (head + mid) / 2
		} else {
			return mid
		}
		// head 等于 mid  说明数组已经遍历完，插入位置即为 head+1
		if head == mid {
			return head + 1
		}
	}
}

func main() {
	nums := []int{1, 3, 5, 6, 8, 9, 10, 11}
	target := 6
	fmt.Println(searchInsert(nums, target))
}
