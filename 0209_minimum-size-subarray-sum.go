package leetcode

import "sort"

/*
  给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的连续子数组，并返回其长度。如果不存在符合条件的连续子数组，返回 0。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/minimum-size-subarray-sum
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func minSubArrayLen(s int, nums []int) int {
	var sum int
	min := 1<<31 - 1
	var l, r int
	for r = range nums {
		sum += nums[r]
		for l < r && sum-nums[l] >= s {
			sum -= nums[l]
			l++
		}
		if sum >= s {
			min = Min(min, r+1-l)
		}
	}
	if min == 1<<31-1 {
		return 0
	}
	return min
}

// Binary Seearch + Prefix Sum
func minSubArrayLen2(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sums := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}
	if sums[len(sums)-1] < s {
		return 0
	}
	min := 1<<31 - 1
	for l := range sums {
		target := s + sums[l]
		if target > sums[len(sums)-1] {
			break
		}
		r := sort.SearchInts(sums, target)
		min = Min(min, r-l)
	}
	if min == 1<<31-1 {
		return 0
	}
	return min
}
