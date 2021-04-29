package main

import (
	"sort"
)

/*
  Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.

  Example 1:
    Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
    Output: [[1,6],[8,10],[15,18]]
    Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].

  Example 2:
    Input: intervals = [[1,4],[4,5]]
    Output: [[1,5]]
    Explanation: Intervals [1,4] and [4,5] are considered overlapping.

  Constraints:
    1. 1 <= intervals.length <= 10^4
    2. intervals[i].length == 2
    3. 0 <= starti <= endi <= 10^4

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/merge-intervals
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	i := 0
	for j := 1; j < len(intervals); j++ {
		if intervals[j][0] > intervals[i][1] {
			i++
			intervals[i] = intervals[j]
		} else if intervals[j][1] > intervals[i][1] {
			intervals[i][1] = intervals[j][1]
		}
	}

	return intervals[:i+1]
}
