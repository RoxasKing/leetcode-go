package codinginterviews

/*
  统计一个数字在排序数组中出现的次数。

  限制：
    0 <= 数组长度 <= 50000
*/

func search(nums []int, target int) int {
	binarySearchFirst := func(l, r int) int {
		for l <= r {
			m := l + (r-l)>>1
			switch {
			case nums[m] > target:
				r = m - 1
			case nums[m] < target:
				l = m + 1
			default:
				if m == 0 || nums[m-1] < target {
					return m
				} else {
					r = m - 1
				}
			}
		}
		return -1
	}
	binarySearchLast := func(l, r int) int {
		for l <= r {
			m := l + (r-l)>>1
			switch {
			case nums[m] > target:
				r = m - 1
			case nums[m] < target:
				l = m + 1
			default:
				if m == len(nums)-1 || nums[m+1] > target {
					return m
				} else {
					l = m + 1
				}
			}
		}
		return -1
	}
	l := binarySearchFirst(0, len(nums)-1)
	if l == -1 {
		return 0
	}
	r := binarySearchLast(0, len(nums)-1)
	return r + 1 - l
}
