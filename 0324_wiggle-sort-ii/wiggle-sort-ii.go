package main

import (
	"math/rand"
)

/*
  给定一个无序的数组 nums，将它重新排列成 nums[0] < nums[1] > nums[2] < nums[3]... 的顺序。

  说明:
    你可以假设所有输入都会得到有效的结果。

  进阶:
    你能用 O(n) 时间复杂度和 / 或原地 O(1) 额外空间来实现吗？
*/

// 3-Way Quick Sort
func wiggleSort(nums []int) {
	m := (len(nums) - 1) >> 1

	threeWayQuickSort(nums, m)

	reverse(nums[:m+1])
	reverse(nums[m+1:])

	n := len(nums)

	tmp := make([]int, n)
	copy(tmp, nums)
	for i := 0; i < n; i += 2 {
		nums[i] = tmp[0]
		tmp = tmp[1:]
	}
	for i := 1; i < n; i += 2 {
		nums[i] = tmp[0]
		tmp = tmp[1:]
	}
}

func reverse(nums []int) {
	for i := 0; i < len(nums)>>1; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
}

func threeWayQuickSort(nums []int, k int) {
	l, r := 0, len(nums)-1
	for l < r {
		i := partition(nums, l, r)
		if i < k {
			l = i + 1
		} else {
			if nums[i] == nums[k] {
				break
			}
			r = i
		}
	}
}

func partition(nums []int, l, r int) int {
	pivotIndex := l + rand.Intn(r+1-l)
	mid := nums[pivotIndex]
	i := l
	for i <= r {
		if nums[i] < mid {
			nums[i], nums[l] = nums[l], nums[i]
			l++
			i++
		} else if nums[i] > mid {
			nums[i], nums[r] = nums[r], nums[i]
			r--
		} else {
			i++
		}
	}
	return i - 1
}
