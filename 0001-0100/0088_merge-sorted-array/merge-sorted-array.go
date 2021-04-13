package main

/*
  Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.

  The number of elements initialized in nums1 and nums2 are m and n respectively. You may assume that nums1 has a size equal to m + n such that it has enough space to hold additional elements from nums2.

  Example 1:
    Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
    Output: [1,2,2,3,5,6]

  Example 2:
    Input: nums1 = [1], m = 1, nums2 = [], n = 0
    Output: [1]

  Constraints:
    1. nums1.length == m + n
    2. nums2.length == n
    3. 0 <= m, n <= 200
    4. 1 <= m + n <= 200
    5. -10^9 <= nums1[i], nums2[i] <= 10^9

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/merge-sorted-array
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
	copy(nums1[:j+1], nums2[:j+1])
}
