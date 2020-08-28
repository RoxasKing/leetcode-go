package main

import "sort"

/*
  给定两个数组，编写一个函数来计算它们的交集。

  说明：
    输出结果中每个元素出现的次数，应与元素在两个数组中出现的次数一致。
    我们可以不考虑输出结果的顺序。

  进阶:
    如果给定的数组已经排好序呢？你将如何优化你的算法？
    如果 nums1 的大小比 nums2 小很多，哪种方法更优？
    如果 nums2 的元素存储在磁盘上，磁盘内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？

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

// Sort
func intersect2(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	var i, j, k int
	for i < len(nums1) && j < len(nums2) {
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
