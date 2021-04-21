package main

import (
	"sort"
)

/*
  给定一个无序的数组，找出数组在排序之后，相邻元素之间最大的差值。
  如果数组元素个数小于 2，则返回 0。

  示例 1:
    输入: [3,6,9,1]
    输出: 3
    解释: 排序后的数组是 [1,3,6,9], 其中相邻元素 (3,6) 和 (6,9) 之间都存在最大差值 3。

  示例 2:
    输入: [10]
    输出: 0
    解释: 数组元素个数小于 2，因此返回 0。

  说明:
    你可以假设数组中所有元素都是非负整数，且数值在 32 位有符号整数范围内。
    请尝试在线性时间复杂度和空间复杂度的条件下解决此问题。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/maximum-gap
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	sort.Ints(nums)
	var max int
	for i := 1; i < len(nums); i++ {
		max = Max(max, nums[i]-nums[i-1])
	}
	return max
}

// Radix Sort
func maximumGap2(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	maxNum := -1 << 31
	for _, num := range nums {
		maxNum = Max(maxNum, num)
	}
	n, buckets := len(nums), 10
	tmps := make([]int, n)
	for exp := 1; maxNum/exp > 0; exp *= 10 {
		count := make([]int, buckets)
		for _, num := range nums {
			count[(num/exp)%10]++
		}
		for i := 1; i < buckets; i++ {
			count[i] += count[i-1]
		}
		for i := n - 1; i >= 0; i-- {
			index := (nums[i] / exp) % 10
			count[index]--
			tmps[count[index]] = nums[i]
		}
		copy(nums, tmps)
	}

	out := 0
	for i := 1; i < n; i++ {
		out = Max(out, nums[i]-nums[i-1])
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
