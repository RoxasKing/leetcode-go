package main

import (
	"sort"
)

/*
  给定一个无序的数组，找出数组在排序之后，相邻元素之间最大的差值。
  如果数组元素个数小于 2，则返回 0。

  说明:
    你可以假设数组中所有元素都是非负整数，且数值在 32 位有符号整数范围内。
    请尝试在线性时间复杂度和空间复杂度的条件下解决此问题。
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
	max := -1<<31 - 1
	for _, num := range nums {
		max = Max(max, num)
	}
	exp, radix := 1, 10
	n := len(nums)
	aux := make([]int, n)
	for max/exp > 0 {
		count := make([]int, radix)
		for i := 0; i < n; i++ {
			count[(nums[i]/exp)%10]++
		}
		for i := 1; i < radix; i++ {
			count[i] += count[i-1]
		}
		for i := n - 1; i >= 0; i-- {
			index := (nums[i] / exp) % 10
			count[index]--
			aux[count[index]] = nums[i]
		}
		for i := 0; i < n; i++ {
			nums[i] = aux[i]
		}
		exp *= 10
	}

	var out int
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
