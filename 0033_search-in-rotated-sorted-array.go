package My_LeetCode_In_Go

/*
  假设按照升序排序的数组在预先未知的某个点上进行了旋转。
  ( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
  搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
  你可以假设数组中不存在重复的元素。
  你的算法时间复杂度必须是 O(log n) 级别。
*/

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	binary_search := func(l, r int) int {
		for l <= r {
			m := (l + r) / 2
			switch {
			case target < nums[m]:
				r = m - 1
			case target > nums[m]:
				l = m + 1
			default:
				return m
			}
		}
		return -1
	}
	var rotate_index int
	l, r := 0, len(nums)-1
	if nums[l] > nums[r] {
		for l <= r {
			m := (l + r) / 2
			if nums[m] > nums[m+1] {
				rotate_index = m + 1
				break
			} else {
				if nums[m] < nums[l] {
					r = m - 1
				} else {
					l = m + 1
				}
			}
		}
	}
	if rotate_index == 0 || target < nums[0] {
		return binary_search(rotate_index, len(nums)-1)
	} else {
		return binary_search(0, rotate_index)
	}
}
