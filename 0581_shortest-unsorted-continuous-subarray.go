package leetcode

import "sort"

/*
  给定一个整数数组，你需要寻找一个连续的子数组，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。
  你找到的子数组应是最短的，请输出它的长度。

  说明 :
    输入的数组长度范围在 [1, 10,000]。
    输入的数组可能包含重复元素 ，所以升序的意思是<=。
*/

// Sort
func findUnsortedSubarray(nums []int) int {
	tmp := make([]int, len(nums))
	copy(tmp, nums)
	sort.Ints(tmp)
	var count int
	for i := 0; i < len(nums); i++ {
		if tmp[i] != nums[i] {
			break
		}
		count++
	}
	if count == len(nums) {
		return 0
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if tmp[i] != nums[i] {
			break
		}
		count++
	}
	return len(nums) - count
}

func findUnsortedSubarray2(nums []int) int {
	max, min := -1<<31, 1<<31-1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			max = Max(max, nums[i])
			min = Min(min, nums[i+1])
		}
	}
	var l, r int
	for i := 0; i < len(nums); i++ {
		if nums[i] > min {
			l = i
			break
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < max {
			r = i
			break
		}
	}
	if l >= r {
		return 0
	}
	return r + 1 - l
}
