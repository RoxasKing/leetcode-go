package My_LeetCode_In_Go

/*
  给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
  你的算法时间复杂度必须是 O(log n) 级别。
  如果数组中不存在目标值，返回 [-1, -1]。
*/

func searchRange(nums []int, target int) []int {
	searchBonde := func(target int, searchLeft bool) int {
		l, r := 0, len(nums)-1
		for l <= r {
			m := (l + r) / 2
			switch {
			case target < nums[m]:
				r = m - 1
			case target > nums[m]:
				l = m + 1
			case target == nums[m]:
				if searchLeft {
					for nums[l] < target {
						l++
					}
					return l
				} else {
					for nums[r] > target {
						r--
					}
					return r
				}
			}
		}
		return -1
	}
	return []int{searchBonde(target, true), searchBonde(target, false)}
}
