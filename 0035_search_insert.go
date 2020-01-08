package My_LeetCode_In_Go

/*
  给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
  你可以假设数组中无重复元素。
*/

func searchInsert(nums []int, target int) int {
	switch {
	case len(nums) == 0 || target < nums[0]:
		return 0
	case target == nums[len(nums)-1]:
		return len(nums) - 1
	case target > nums[len(nums)-1]:
		return len(nums)
	}
	l, r := 0, len(nums)-1
	if target > nums[r] {
		return r + 1
	}
	for l < r {
		m := (l + r) / 2
		switch {
		case target == nums[m]:
			return m
		case target < nums[m]:
			r = m
		case target > nums[m]:
			l = m + 1
		}
	}
	return l
}
