package main

/*
  Given an array of integers nums and an integer k, return the total number of continuous subarrays whose sum equals to k.

  Example 1:
    Input: nums = [1,1,1], k = 2
    Output: 2

  Example 2:
    Input: nums = [1,2,3], k = 3
    Output: 2

  Constraints:
    1. 1 <= nums.length <= 2 * 10^4
    2. -1000 <= nums[i] <= 1000
    3. -10^7 <= k <= 10^7

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/subarray-sum-equals-k
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Prefix Sum + Hash
func subarraySum(nums []int, k int) int {
	out := 0
	sum := 0
	cnt := map[int]int{0: 1} // if sum-k == 0 , count++
	for _, num := range nums {
		sum += num
		if cnt[sum-k] > 0 {
			out += cnt[sum-k]
		}
		cnt[sum]++
	}
	return out
}
