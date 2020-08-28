package main

/*
  给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
  你的算法时间复杂度必须是 O(log n) 级别。
  如果数组中不存在目标值，返回 [-1, -1]。
*/

func searchRange(nums []int, target int) []int {
	out := []int{-1, -1}
	out[0] = searchBorder(nums, 0, len(nums)-1, target, true)
	if out[0] == -1 {
		return out
	}
	out[1] = searchBorder(nums, out[0], len(nums)-1, target, false)
	return out
}

func searchBorder(nums []int, l, r, target int, searchLeft bool) int {
	for l <= r {
		m := l + (r-l)>>1
		if nums[m] < target {
			l = m + 1
		} else if nums[m] > target {
			r = m - 1
		} else if searchLeft {
			if m == 0 || target != nums[m-1] {
				return m
			}
			r = m - 1
		} else {
			if m == len(nums)-1 || target != nums[m+1] {
				return m
			}
			l = m + 1
		}
	}
	return -1
}
