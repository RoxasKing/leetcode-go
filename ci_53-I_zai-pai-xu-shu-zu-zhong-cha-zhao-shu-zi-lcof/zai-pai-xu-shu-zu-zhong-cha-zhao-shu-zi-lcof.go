package main

/*
  统计一个数字在排序数组中出现的次数。

  示例 1:
    输入: nums = [5,7,7,8,8,10], target = 8
    输出: 2

  示例 2:
    输入: nums = [5,7,7,8,8,10], target = 6
    输出: 0

  限制：
    0 <= 数组长度 <= 50000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Binary Search
func search(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1
	start := -1
	for l <= r {
		m := (l + r) >> 1
		if nums[m] < target {
			l = m + 1
		} else if nums[m] > target {
			r = m - 1
		} else {
			if m == 0 || nums[m-1] < target {
				start = m
				break
			} else {
				r = m - 1
			}
		}
	}
	if start == -1 {
		return 0
	}

	l, r = start, n-1
	end := -1
	for l <= r {
		m := (l + r) >> 1
		if nums[m] < target {
			l = m + 1
		} else if nums[m] > target {
			r = m - 1
		} else {
			if m == n-1 || nums[m+1] > target {
				end = m
				break
			} else {
				l = m + 1
			}
		}
	}
	return end + 1 - start
}
