package main

/*
  给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
  你的算法时间复杂度必须是 O(log n) 级别。
  如果数组中不存在目标值，返回 [-1, -1]。

  示例 1:
    输入: nums = [5,7,7,8,8,10], target = 8
    输出: [3,4]

  示例 2:
    输入: nums = [5,7,7,8,8,10], target = 6
    输出: [-1,-1]

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Binary Search
func searchRange(nums []int, target int) []int {
	n := len(nums)
	l := bSearch(nums, target, 0, n-1, true)
	if l == -1 {
		return []int{-1, -1}
	}
	r := bSearch(nums, target, l, n-1, false)
	return []int{l, r}
}

func bSearch(nums []int, target, l, r int, leftBond bool) int {
	for l <= r {
		m := l + (r-l)>>1
		if nums[m] < target {
			l = m + 1
		} else if nums[m] > target {
			r = m - 1
		} else if leftBond {
			if m == 0 || nums[m-1] < target {
				return m
			}
			r = m - 1
		} else {
			if m == r || nums[m+1] > target {
				return m
			}
			l = m + 1
		}
	}
	return -1
}
