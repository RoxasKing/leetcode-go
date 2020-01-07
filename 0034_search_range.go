package My_LeetCode_In_Go

/*
  给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
  你的算法时间复杂度必须是 O(log n) 级别。
  如果数组中不存在目标值，返回 [-1, -1]。
*/

func searchRange(nums []int, target int) []int {
	if len(nums) < 1 {
		return []int{-1, -1}
	}
	searchBonde := func(target int, searchLeft bool) int {
		left, right := 0, len(nums)
		for left < right {
			mid := (left + right) / 2
			if target < nums[mid] || (searchLeft && target == nums[mid]) {
				right = mid
			} else {
				left = mid + 1
			}
		}
		return left
	}
	l := searchBonde(target, true)
	if l == len(nums) || target != nums[l] {
		return []int{-1, -1}
	}
	r := searchBonde(target, false) - 1
	return []int{l, r}
}
