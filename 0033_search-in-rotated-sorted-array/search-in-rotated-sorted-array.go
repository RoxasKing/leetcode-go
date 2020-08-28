package main

/*
  假设按照升序排序的数组在预先未知的某个点上进行了旋转。
  ( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
  搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
  你可以假设数组中不存在重复的元素。
  你的算法时间复杂度必须是 O(log n) 级别。
*/

// Binary Search
func search(nums []int, target int) int {
	rotateIndex := binarySearchRotateIndex(nums)
	if rotateIndex == 0 || nums[0] > target {
		return binarySearchTargetIndex(nums, rotateIndex, len(nums)-1, target)
	}
	return binarySearchTargetIndex(nums, 0, rotateIndex, target)
}

func binarySearchRotateIndex(nums []int) int {
	l, r := 0, len(nums)-1
	if nums[l] > nums[r] {
		for l < r {
			m := l + (r-l)>>1
			if nums[m] > nums[m+1] {
				return m + 1
			}
			if nums[m] < nums[l] {
				r = m
			} else {
				l = m + 1
			}
		}
	}
	return 0
}

func binarySearchTargetIndex(nums []int, l, r, target int) int {
	for l <= r {
		m := l + (r-l)>>1
		if nums[m] < target {
			l = m + 1
		} else if nums[m] > target {
			r = m - 1
		} else {
			return m
		}
	}
	return -1
}
