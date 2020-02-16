package leetcode

/*
  给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
  你可以假设数组中无重复元素。
*/

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)>>1
		switch {
		case nums[m] < target:
			l = m + 1
		case nums[m] > target:
			r = m - 1
		default:
			return m
		}
	}
	return l
}
