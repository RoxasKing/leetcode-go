package main

import "math/rand"

/*
  给定一个整数数组 nums，将该数组升序排列。
*/

// Merge Sort
func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	return mergeSortArray(sortArray(nums[:mid]), sortArray(nums[mid:]))
}

func mergeSortArray(l, r []int) []int {
	out := make([]int, 0, len(l)+len(r))
	for len(l) != 0 && len(r) != 0 {
		if l[0] < r[0] {
			out = append(out, l[0])
			l = l[1:]
		} else {
			out = append(out, r[0])
			r = r[1:]
		}
	}
	if len(l) != 0 {
		out = append(out, l...)
	}
	if len(r) != 0 {
		out = append(out, r...)
	}
	return out
}

// Quick Sort
func sortArray2(nums []int) []int {
	quickSort(0, len(nums)-1, nums)
	return nums
}

func quickSort(l, r int, nums []int) {
	if l >= r {
		return
	}
	pivotIdx := rand.Intn(r-l+1) + l
	nums[pivotIdx], nums[r] = nums[r], nums[pivotIdx]
	pivot, mid := nums[r], l
	for i := l; i < r; i++ {
		if nums[i] < pivot {
			nums[i], nums[mid] = nums[mid], nums[i]
			mid++
		}
	}
	nums[mid], nums[r] = nums[r], nums[mid]
	quickSort(l, mid-1, nums)
	quickSort(mid+1, r, nums)
}
