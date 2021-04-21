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
	n := len(nums)
	k := (n - 1) >> 1

	threeWayQuickSort(nums, n, k)

	reverse(nums[:k+1], n)
	reverse(nums[k+1:], n)

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

func reverse(nums []int, n int) {
	for i := 0; i < n>>1; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
}

func threeWayQuickSort(nums []int, n, k int) {
	l, r := 0, n-1
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
	pivot := nums[pivotIndex]
	for i := l; i <= r; {
		if nums[i] < pivot {
			nums[i], nums[l] = nums[l], nums[i]
			l++
			i++
		} else if nums[i] > pivot {
			nums[i], nums[r] = nums[r], nums[i]
			r--
		} else {
			i++
		}
	}
	return r
}
