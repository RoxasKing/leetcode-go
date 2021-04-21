package main

/*
  You are given two arrays of integers nums1 and nums2, possibly of different lengths. The values in the arrays are between 1 and 6, inclusive.

  In one operation, you can change any integer's value in any of the arrays to any value between 1 and 6, inclusive.

  Return the minimum number of operations required to make the sum of values in nums1 equal to the sum of values in nums2. Return -1​​​​​ if it is not possible to make the sum of the two arrays equal.

  Example 1:
    Input: nums1 = [1,2,3,4,5,6], nums2 = [1,1,2,2,2,2]
    Output: 3
    Explanation: You can make the sums of nums1 and nums2 equal with 3 operations. All indices are 0-indexed.
    - Change nums2[0] to 6. nums1 = [1,2,3,4,5,6], nums2 = [6,1,2,2,2,2].
    - Change nums1[5] to 1. nums1 = [1,2,3,4,5,1], nums2 = [6,1,2,2,2,2].
    - Change nums1[2] to 2. nums1 = [1,2,2,4,5,1], nums2 = [6,1,2,2,2,2].

  Example 2:
    Input: nums1 = [1,1,1,1,1,1,1], nums2 = [6]
    Output: -1
    Explanation: There is no way to decrease the sum of nums1 or to increase the sum of nums2 to make them equal.

  Example 3:
    Input: nums1 = [6,6], nums2 = [1]
    Output: 3
    Explanation: You can make the sums of nums1 and nums2 equal with 3 operations. All indices are 0-indexed.
    - Change nums1[0] to 2. nums1 = [2,6], nums2 = [1].
    - Change nums1[1] to 2. nums1 = [2,2], nums2 = [1].
    - Change nums2[0] to 4. nums1 = [2,2], nums2 = [4].

  Constraints:
    1. 1 <= nums1.length, nums2.length <= 10^5
    2. 1 <= nums1[i], nums2[i] <= 6

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/equal-sum-arrays-with-minimum-number-of-operations
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Greedy Algrorithm
func minOperations(nums1 []int, nums2 []int) int {
	sum1, sum2 := 0, 0
	cnt1, cnt2 := [7]int{}, [7]int{}
	for _, num := range nums1 {
		sum1 += num
		cnt1[num]++
	}
	for _, num := range nums2 {
		sum2 += num
		cnt2[num]++
	}

	out := 0
	for sum1 != sum2 {
		out++

		if sum1 < sum2 {
			sum1, sum2 = sum2, sum1
			cnt1, cnt2 = cnt2, cnt1
		}
		diff := sum1 - sum2

		dec, decIdx := -1, -1
		for i := 6; i > 1; i-- {
			if cnt1[i] > 0 {
				dec, decIdx = i-1, i
				if i-diff > 1 {
					dec = diff
				}
				break
			}
		}

		if dec == diff {
			break
		}

		inc, incIdx := -1, -1
		for i := 1; i < 6; i++ {
			if cnt2[i] > 0 {
				inc, incIdx = 6-i, i
				if i+diff < 6 {
					inc = diff
				}
				break
			}
		}

		if inc == diff {
			break
		}

		if decIdx == -1 && incIdx == -1 {
			return -1
		} else if decIdx != -1 && incIdx != -1 {
			if dec > inc {
				sum1 -= dec
				cnt1[decIdx]--
				cnt1[decIdx-dec]++
			} else {
				sum2 += inc
				cnt2[incIdx]--
				cnt2[incIdx+inc]++
			}
		} else if decIdx != -1 {
			sum1 -= dec
			cnt1[decIdx]--
			cnt1[decIdx-dec]++
		} else if incIdx != -1 {
			sum2 += inc
			cnt2[incIdx]--
			cnt2[incIdx+inc]++
		}

	}
	return out
}
