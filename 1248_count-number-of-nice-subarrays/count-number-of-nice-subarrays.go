package main

/*
  Given an array of integers nums and an integer k. A continuous subarray is called nice if there are k odd numbers on it.
  Return the number of nice sub-arrays.

  Example 1:
    Input: nums = [1,1,2,1,1], k = 3
    Output: 2
    Explanation: The only sub-arrays with 3 odd numbers are [1,1,2,1] and [1,2,1,1].

  Example 2:
    Input: nums = [2,4,6], k = 1
    Output: 0
    Explanation: There is no odd numbers in the array.

  Example 3:
    Input: nums = [2,2,2,1,2,2,1,2,2,2], k = 2
    Output: 16

  Constraints:
    1. 1 <= nums.length <= 50000
    2. 1 <= nums[i] <= 10^5
    3. 1 <= k <= nums.length

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/count-number-of-nice-subarrays
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Sliding Window
func numberOfSubarrays(nums []int, k int) int {
	oddIdx := []int{-1}
	for i := range nums {
		if nums[i]&1 == 1 {
			oddIdx = append(oddIdx, i)
		}
	}
	oddIdx = append(oddIdx, len(nums))

	if len(oddIdx)-2 < k {
		return 0
	}

	out := 0
	for i := 1; i <= len(oddIdx)-1-k; i++ {
		lCnt := oddIdx[i] - (oddIdx[i-1] + 1)
		rCnt := oddIdx[i+k] - 1 - oddIdx[i+k-1]
		out += 1 + lCnt + rCnt + lCnt*rCnt
	}
	return out
}

// Two Pointers + Sliding Window
func numberOfSubarrays2(nums []int, k int) int {
	odd := 0
	out := 0
	for l, r := 0, 0; r < len(nums); r++ {
		if nums[r]&1 == 1 {
			odd++
		}
		for l <= r && odd > k {
			if nums[l]&1 == 1 {
				odd--
			}
			l++
		}
		if odd == k {
			for i := l; i <= r; i++ {
				out++
				if nums[i]&1 == 1 {
					break
				}
			}
		}
	}
	return out
}
