package leetcode

import "sort"

/*
  给定两个数组，编写一个函数来计算它们的交集。

  说明：
    输出结果中的每个元素一定是唯一的。
    我们可以不考虑输出结果的顺序。
*/

// Binary Search
func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) < len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	sort.Ints(nums1)
	sort.Ints(nums2)
	binarySearch := func(target int) int {
		l, r := 0, len(nums1)-1
		for l <= r {
			m := l + (r-l)>>1
			if nums1[m] < target {
				l = m + 1
			} else if nums1[m] > target {
				r = m - 1
			} else {
				return m
			}
		}
		return -1
	}
	var out []int
	for i := 0; i < len(nums2); i++ {
		if index := binarySearch(nums2[i]); index != -1 {
			out = append(out, nums1[index])
		}
		for i+1 < len(nums2) && nums2[i+1] == nums2[i] {
			i++
		}
	}
	return out
}

// Hash
func intersection2(nums1, nums2 []int) []int {
	sets1 := make(map[int]struct{})
	sets2 := make(map[int]struct{})
	for _, num := range nums1 {
		sets1[num] = struct{}{}
	}
	for _, num := range nums2 {
		sets2[num] = struct{}{}
	}
	if len(sets1) < len(sets2) {
		sets1, sets2 = sets2, sets1
	}
	var out []int
	for num := range sets2 {
		if _, ok := sets1[num]; ok {
			out = append(out, num)
		}
	}
	return out
}
