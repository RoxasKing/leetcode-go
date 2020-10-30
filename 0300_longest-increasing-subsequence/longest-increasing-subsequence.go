package main

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

// Dynamic Programming
func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	max := dp[0]
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[j] < nums[i] {
				dp[i] = Max(dp[i], dp[j]+1)
			}
		}
		max = Max(max, dp[i])
	}
	return max
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// Binary Search
func lengthOfLIS2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	mins := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		if nums[i] > mins[len(mins)-1] {
			mins = append(mins, nums[i])
		} else if nums[i] < mins[len(mins)-1] {
			mins[binarySearch(mins, nums[i])] = nums[i]
		}
	}
	return len(mins)
}

// search the specify number's index in the array
func binarySearch(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := l + (r-l)>>1
		if nums[m] < target {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
