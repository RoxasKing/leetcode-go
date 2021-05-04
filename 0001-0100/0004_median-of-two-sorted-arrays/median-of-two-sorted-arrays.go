package main

/*
  Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

  Example 1:
    Input: nums1 = [1,3], nums2 = [2]
    Output: 2.00000
    Explanation: merged array = [1,2,3] and median is 2.

  Example 2:
    Input: nums1 = [1,2], nums2 = [3,4]
    Output: 2.50000
    Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.

  Example 3:
    Input: nums1 = [0,0], nums2 = [0,0]
    Output: 0.00000

  Example 4:
    Input: nums1 = [], nums2 = [1]
    Output: 1.00000

  Example 5:
    Input: nums1 = [2], nums2 = []
    Output: 2.00000

  Constraints:
    1. nums1.length == m
    2. nums2.length == n
    3. 0 <= m <= 1000
    4. 0 <= n <= 1000
    5. 1 <= m + n <= 2000
    6. -10^6 <= nums1[i], nums2[i] <= 10^6

  Follow up: The overall run time complexity should be O(log (m+n)).

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Binary Search
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	n1, n2 := len(nums1), len(nums2)

	l1, r1 := 0, n1
	m, m1, m2 := (n1+n2+1)>>1, 0, 0
	for l1 <= r1 {
		m1 = (l1 + r1) >> 1
		m2 = m - m1
		if m1 < r1 && nums1[m1] < nums2[m2-1] {
			l1 = m1 + 1
		} else if m1 > l1 && nums1[m1-1] > nums2[m2] {
			r1 = m1 - 1
		} else {
			break
		}
	}

	if n1 == n2 {
		if m1 == 0 {
			return float64(nums2[n2-1]+nums1[0]) / 2
		} else if m2 == 0 {
			return float64(nums1[n1-1]+nums2[0]) / 2
		}
		return float64(Max(nums1[m1-1], nums2[m2-1])+Min(nums1[m1], nums2[m2])) / 2
	}

	maxL := nums2[m2-1]
	if m1 > 0 {
		maxL = Max(maxL, nums1[m1-1])
	}
	if (n1+n2)&1 == 1 {
		return float64(maxL)
	}
	minR := nums2[m2]
	if m1 < n1 {
		minR = Min(minR, nums1[m1])
	}
	return float64(maxL+minR) / 2
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
