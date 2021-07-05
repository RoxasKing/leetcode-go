package main

import "sort"

// Hash
func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	count := make(map[int]int)
	for _, num := range nums1 {
		count[num]++
	}
	var out []int
	for _, num := range nums2 {
		if count[num] > 0 {
			count[num]--
			out = append(out, num)
		}
	}
	return out
}

// Two Pointers
func intersect2(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	m, n := len(nums1), len(nums2)
	i, j, k := 0, 0, 0
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			nums1[k] = nums1[i]
			i++
			j++
			k++
		}
	}
	return nums1[:k]
}
