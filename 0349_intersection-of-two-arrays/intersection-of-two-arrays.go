package main

import "sort"

/*
  给定两个数组，编写一个函数来计算它们的交集。

  示例 1：
    输入：nums1 = [1,2,2,1], nums2 = [2,2]
    输出：[2]

  示例 2：
    输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
    输出：[9,4]

  说明：
    输出结果中的每个元素一定是唯一的。
    我们可以不考虑输出结果的顺序。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/intersection-of-two-arrays
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

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
