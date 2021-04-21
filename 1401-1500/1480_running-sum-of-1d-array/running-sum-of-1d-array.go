package main

/*
  Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).

  Return the running sum of nums.

  Example 1:
    Input: nums = [1,2,3,4]
    Output: [1,3,6,10]
    Explanation: Running sum is obtained as follows: [1, 1+2, 1+2+3, 1+2+3+4].

  Example 2:
    Input: nums = [1,1,1,1,1]
    Output: [1,2,3,4,5]
    Explanation: Running sum is obtained as follows: [1, 1+1, 1+1+1, 1+1+1+1, 1+1+1+1+1].

  Example 3:
    Input: nums = [3,1,2,10,1]
    Output: [3,4,6,16,17]

  Constraints:
    1. 1 <= nums.length <= 1000
    2. -10^6 <= nums[i] <= 10^6

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/running-sum-of-1d-array
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func runningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return nums
}
