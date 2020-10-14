package main

/*
  给定一个只包含正整数的非空数组。是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

  注意:
    每个数组中的元素不会超过 100
    数组的大小不会超过 200

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/partition-equal-subset-sum
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Dynamic Programming
func canPartition(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}

	sum, max := 0, -1<<31
	for _, num := range nums {
		sum += num
		max = Max(max, num)
	}
	target := sum >> 1
	if sum&1 == 1 || max > target {
		return false
	}

	dp := make([]bool, target+1)
	dp[0] = true

	for i := 1; i < n; i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}

	return dp[target]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
