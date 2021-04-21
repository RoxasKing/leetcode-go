package main

import "sort"

/*
  Given two integer arrays nums1 and nums2, return an array of their intersection. Each element in the result must appear as many times as it shows in both arrays and you may return the result in any order.

  Example 1:
    Input: nums1 = [1,2,2,1], nums2 = [2,2]
    Output: [2,2]

  Example 2:
    Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
    Output: [4,9]
    Explanation: [9,4] is also accepted.

  Constraints:
    1. 1 <= nums1.length, nums2.length <= 1000
    2. 0 <= nums1[i], nums2[i] <= 1000

  Follow up:
    1. What if the given array is already sorted? How would you optimize your algorithm?
    2. What if nums1's size is small compared to nums2's size? Which algorithm is better?
    3. What if elements of nums2 are stored on disk, and the memory is limited such that you cannot load all elements into the memory at once?

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/intersection-of-two-arrays-ii
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

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
