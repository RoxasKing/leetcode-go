package main

import "sort"

// Hash
func intersection(nums1, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	dict := make(map[int]bool)
	for _, num := range nums1 {
		dict[num] = true
	}
	out := []int{}
	for _, num := range nums2 {
		if dict[num] {
			out = append(out, num)
			delete(dict, num)
		}
	}
	return out
}

// Binary Search
func intersection2(nums1 []int, nums2 []int) []int {
	if len(nums1) < len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	sort.Ints(nums1)
	sort.Ints(nums2)
	out := []int{}
	for i := 0; i < len(nums2); i++ {
		if index := binarySearch(nums1, nums2[i]); index != -1 {
			out = append(out, nums1[index])
		}
		for i+1 < len(nums2) && nums2[i+1] == nums2[i] {
			i++
		}
	}
	return out
}

func binarySearch(nums []int, target int) int {
	l, r := 0, len(nums)-1
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
