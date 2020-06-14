package leetcode

import "sort"

/*
  给你一个整数数组 arr 和一个目标值 target ，请你返回一个整数 value ，使得将数组中所有大于 value 的值变成 value 后，数组的和最接近  target （最接近表示两者之差的绝对值最小）。

  如果有多种使得和最接近 target 的方案，请你返回这些整数中的最小值。

  请注意，答案不一定是 arr 中的数字。

  提示：
    1 <= arr.length <= 10^4
    1 <= arr[i], target <= 10^5

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/sum-of-mutated-array-closest-to-target
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Binary Search
func findBestValue(arr []int, target int) int {
	sort.Ints(arr)
	prefix := make([]int, len(arr)+1)
	for i := 1; i < len(prefix); i++ {
		prefix[i] = prefix[i-1] + arr[i-1]
	}
	out, diff := 1<<31-1, target
	l, r := 0, arr[len(arr)-1]
	for l <= r {
		m := l + (r-l)>>1
		// index := sort.SearchInts(arr, m)
		index := findFirstNotLessIndex(arr, m)
		cur := prefix[index] + m*(len(arr)-index)
		newDiff := Abs(cur - target)
		if newDiff < diff {
			out, diff = m, newDiff
		} else if newDiff == diff {
			out = Min(out, m)
		}
		if cur < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return out
}

func findFirstNotLessIndex(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)>>1
		if nums[m] >= target {
			if m == 0 || nums[m-1] < target {
				return m
			}
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
}
