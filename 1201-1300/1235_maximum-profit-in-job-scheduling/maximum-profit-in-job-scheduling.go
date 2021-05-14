package main

import "sort"

/*
  We have n jobs, where every job is scheduled to be done from startTime[i] to endTime[i], obtaining a profit of profit[i].

  You're given the startTime, endTime and profit arrays, return the maximum profit you can take such that there are no two jobs in the subset with overlapping time range.

  If you choose a job that ends at time X you will be able to start another job that starts at time X.

  Example 1:
    Input: startTime = [1,2,3,3], endTime = [3,4,5,6], profit = [50,10,40,70]
    Output: 120
    Explanation: The subset chosen is the first and fourth job.
      Time range [1-3]+[3-6] , we get profit of 120 = 50 + 70.

  Example 2:
    Input: startTime = [1,2,3,4,6], endTime = [3,5,10,6,9], profit = [20,20,100,70,60]
    Output: 150
    Explanation: The subset chosen is the first, fourth and fifth job.
      Profit obtained 150 = 20 + 70 + 60.

  Example 3:
    Input: startTime = [1,1,1], endTime = [2,3,4], profit = [5,6,4]
    Output: 6

  Constraints:
    1. 1 <= startTime.length == endTime.length == profit.length <= 5 * 10^4
    2. 1 <= startTime[i] < endTime[i] <= 10^9
    3. 1 <= profit[i] <= 10^4

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/maximum-profit-in-job-scheduling
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Dynamic Programming

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	n := len(startTime)
	set := map[int]struct{}{}
	for i := 0; i < n; i++ {
		set[startTime[i]] = struct{}{}
		set[endTime[i]] = struct{}{}
	}

	times := make([]int, 0, len(set))
	for t := range set {
		times = append(times, t)
	}
	sort.Ints(times)

	dict := map[int]int{}
	for i, point := range times {
		dict[point] = i
	}

	jobs := make([][3]int, n)
	for i := 0; i < n; i++ {
		jobs[i] = [3]int{startTime[i], endTime[i], profit[i]}
	}
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i][1] < jobs[j][1]
	})

	m := len(times)
	f := make([]int, m)

	for i := 1; i < m; i++ {
		f[i] = f[i-1]
		for len(jobs) > 0 && jobs[0][1] == times[i] {
			job := jobs[0]
			jobs = jobs[1:]
			f[i] = Max(f[i], f[dict[job[0]]]+job[2])
		}
	}

	return f[m-1]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
