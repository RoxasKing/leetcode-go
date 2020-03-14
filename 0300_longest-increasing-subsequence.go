package leetcode

/*
  给定一个无序的整数数组，找到其中最长上升子序列的长度。
  示例:
    输入: [10,9,2,5,3,7,101,18]
    输出: 4
    解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
  说明:
    可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
    你算法的时间复杂度应该为 O(n2) 。
  进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?
*/

func lengthOfLIS(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	binarySearch := func(nums []int, target int) int {
		l, r := 0, len(nums)-1
		for l <= r {
			m := l + (r-l)>>1
			if nums[m] >= target {
				if m == 0 || nums[m-1] < target {
					return m
				}
				r = m - 1
			} else {
				l = m + 1
			}
		}
		return -1
	}
	mins := append(make([]int, 0, len(nums)), nums[1])
	for i := 1; i < len(nums); i++ {
		if nums[i] > mins[len(mins)-1] {
			mins = append(mins, nums[i])
		} else {
			mins[binarySearch(mins, nums[i])] = nums[i]
		}
	}
	return len(mins)
}

// Dynamic Programming
func lengthOfLIS2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	max := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = Max(dp[j]+1, dp[i])
			}
		}
		max = Max(max, dp[i])
	}
	return max
}
