package main

/*
  Given a non-empty array nums containing only positive integers, find if the array can be partitioned into two subsets such that the sum of elements in both subsets is equal.

  Example 1:
    Input: nums = [1,5,11,5]
    Output: true
    Explanation: The array can be partitioned as [1, 5, 5] and [11].

  Example 2:
    Input: nums = [1,2,3,5]
    Output: false
    Explanation: The array cannot be partitioned into equal sum subsets.

  Constraints:
    1. 1 <= nums.length <= 200
    2. 1 <= nums[i] <= 100

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/partition-equal-subset-sum
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Dynamic Programming

func canPartition(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum&1 == 1 {
		return false
	}

	m := sum >> 1
	dp := make([]bool, m+1)
	dp[0] = true

	for i := 0; i < n; i++ {
		for j := m; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}

	return dp[m]
}
