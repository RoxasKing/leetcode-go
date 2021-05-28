package main

/*
  The Hamming distance between two integers is the number of positions at which the corresponding bits are different.

  Given an integer array nums, return the sum of Hamming distances between all the pairs of the integers in nums.

  Example 1:
    Input: nums = [4,14,2]
    Output: 6
    Explanation: In binary representation, the 4 is 0100, 14 is 1110, and 2 is 0010 (just
      showing the four bits relevant in this case).
      The answer will be:
      HammingDistance(4, 14) + HammingDistance(4, 2) + HammingDistance(14, 2) = 2 + 2 + 2 = 6.

  Example 2:
    Input: nums = [4,14,4]
    Output: 4

  Constraints:
    1. 1 <= nums.length <= 10^5
    2. 0 <= nums[i] <= 10^9

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/total-hamming-distance
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Bit Manipulation

func totalHammingDistance(nums []int) int {
	out := 0
	cnt := [32][2]int{}
	for _, num := range nums {
		for i := 31; i >= 0; i-- {
			idx := (num >> i) & 1
			out += cnt[i][1-idx]
			cnt[i][idx]++
		}
	}
	return out
}
