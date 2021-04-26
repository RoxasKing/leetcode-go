package main

import (
	"sort"
)

/*
  You are given an array of intervals, where intervals[i] = [starti, endi] and each starti is unique.

  The right interval for an interval i is an interval j such that startj >= endi and startj is minimized.

  Return an array of right interval indices for each interval i. If no right interval exists for interval i, then put -1 at index i.

  Example 1:
    Input: intervals = [[1,2]]
    Output: [-1]
    Explanation: There is only one interval in the collection, so it outputs -1.

  Example 2:
    Input: intervals = [[3,4],[2,3],[1,2]]
    Output: [-1,0,1]
    Explanation: There is no right interval for [3,4].
      The right interval for [2,3] is [3,4] since start0 = 3 is the smallest start that is >= end1 = 3.
      The right interval for [1,2] is [2,3] since start1 = 2 is the smallest start that is >= end2 = 2.

  Example 3:
    Input: intervals = [[1,4],[2,3],[3,4]]
    Output: [-1,2,-1]
    Explanation: There is no right interval for [1,4] and [3,4].
      The right interval for [2,3] is [3,4] since start2 = 3 is the smallest start that is >= end1 = 3.

  Constraints:
    1. 1 <= intervals.length <= 2 * 10^4
    2. intervals[i].length == 2
    3. -10^6 <= starti <= endi <= 10^6
    4. The start point of each interval is unique.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/find-right-interval
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Binary Search
func findRightInterval(intervals [][]int) []int {
	n := len(intervals)
	idxs := make([]int, n)
	for i := 0; i < n; i++ {
		idxs[i] = i
	}
	sort.Slice(idxs, func(i, j int) bool {
		return intervals[idxs[i]][0] < intervals[idxs[j]][0]
	})

	out := make([]int, n)
	for i, interval := range intervals {
		j := sort.Search(n, func(j int) bool {
			return intervals[idxs[j]][0] >= interval[1]
		})

		if j == n {
			out[i] = -1
		} else {
			out[i] = idxs[j]
		}
	}
	return out
}
