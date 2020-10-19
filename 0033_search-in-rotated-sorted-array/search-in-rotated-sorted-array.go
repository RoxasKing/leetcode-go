package main

/*
  给你一个升序排列的整数数组 nums ，和一个整数 target 。
  假设按照升序排序的数组在预先未知的某个点上进行了旋转。（例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] ）。
  请你在数组中搜索 target ，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。

  示例 1：
    输入：nums = [4,5,6,7,0,1,2], target = 0
    输出：4

  示例 2：
    输入：nums = [4,5,6,7,0,1,2], target = 3
    输出：-1

  示例 3：
    输入：nums = [1], target = 0
    输出：-1

  提示：
    1 <= nums.length <= 5000
    -10^4 <= nums[i] <= 10^4
    nums 中的每个值都 独一无二
    nums 肯定会在某个点上旋转
    -10^4 <= target <= 10^4

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/search-in-rotated-sorted-array
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Binary Search
func search(nums []int, target int) int {
	n := len(nums)
	rotateIndex := binarySearchRotateIndex(nums, 0, n-1)
	l, r := 0, n-1
	if rotateIndex == 0 || nums[0] > target {
		l = rotateIndex
	} else {
		r = rotateIndex
	}
	return binarySearchTargetIndex(nums, l, r, target)
}

func binarySearchRotateIndex(nums []int, l, r int) int {
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
