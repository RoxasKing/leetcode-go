package main

import "math/bits"

/*
  You are given two integer arrays nums1 and nums2 of length n.

  The XOR sum of the two integer arrays is (nums1[0] XOR nums2[0]) + (nums1[1] XOR nums2[1]) + ... + (nums1[n - 1] XOR nums2[n - 1]) (0-indexed).

    For example, the XOR sum of [1,2,3] and [3,2,1] is equal to (1 XOR 3) + (2 XOR 2) + (3 XOR 1) = 2 + 0 + 2 = 4.

  Rearrange the elements of nums2 such that the resulting XOR sum is minimized.

  Return the XOR sum after the rearrangement.

  Example 1:
    Input: nums1 = [1,2], nums2 = [2,3]
    Output: 2
    Explanation: Rearrange nums2 so that it becomes [3,2].
      The XOR sum is (1 XOR 3) + (2 XOR 2) = 2 + 0 = 2.

  Example 2:
    Input: nums1 = [1,0,3], nums2 = [5,3,4]
    Output: 8
    Explanation: Rearrange nums2 so that it becomes [5,4,3].
      The XOR sum is (1 XOR 5) + (0 XOR 4) + (3 XOR 3) = 4 + 4 + 0 = 8.


  Constraints:
    1. n == nums1.length
    2. n == nums2.length
    3. 1 <= n <= 14
    4. 0 <= nums1[i], nums2[i] <= 10^7

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/minimum-xor-sum-of-two-arrays
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Dynamic Programming
// Bit Manipulation

func minimumXORSum(nums1 []int, nums2 []int) int {
	n := len(nums1)
	f := make([]int, 1<<n)
	for i := range f {
		f[i] = int(1e9)
	}
	f[0] = 0

	for mask := 1; mask < 1<<n; mask++ {
		c := bits.OnesCount(uint(mask))
		for i := 0; i < n; i++ {
			if mask&(1<<i) > 0 {
				f[mask] = Min(f[mask], f[mask^(1<<i)]+(nums1[c-1]^nums2[i]))
			}
		}
	}

	return f[(1<<n)-1]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
