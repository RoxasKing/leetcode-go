package main

import "math/bits"

/*
  You are given an integer array jobs, where jobs[i] is the amount of time it takes to complete the ith job.

  There are k workers that you can assign jobs to. Each job should be assigned to exactly one worker. The working time of a worker is the sum of the time it takes to complete all jobs assigned to them. Your goal is to devise an optimal assignment such that the maximum working time of any worker is minimized.

  Return the minimum possible maximum working time of any assignment.

  Example 1:
    Input: jobs = [3,2,3], k = 3
    Output: 3
    Explanation: By assigning each person one job, the maximum time is 3.

  Example 2:
    Input: jobs = [1,2,4,7,8], k = 2
    Output: 11
    Explanation: Assign the jobs the following way:
      Worker 1: 1, 2, 8 (working time = 1 + 2 + 8 = 11)
      Worker 2: 4, 7 (working time = 4 + 7 = 11)
      The maximum working time is 11.

  Constraints:
    1. 1 <= k <= jobs.length <= 12
    2. 1 <= jobs[i] <= 10^7

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/find-minimum-time-to-finish-all-jobs
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Dynamic Programming
func minimumTimeRequired(jobs []int, k int) int {
	n := len(jobs)
	sum := make([]int, 1<<n)
	for i := 1; i < 1<<n; i++ {
		j := bits.TrailingZeros(uint(i))
		sum[i] = sum[i^1<<j] + jobs[j]
	}

	f0 := make([]int, 1<<n)
	f1 := make([]int, 1<<n)
	for i := 0; i < 1<<n; i++ {
		f0[i] = sum[i]
	}

	for ; k > 1; k-- {
		for i := 1; i < 1<<n; i++ {
			f1[i] = 1<<63 - 1
			for j := i; j > 0; j = (j - 1) & i {
				f1[i] = Min(f1[i], Max(f0[i-j], sum[j]))
			}
		}
		f0, f1 = f1, f0
	}

	return f0[1<<n-1]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
