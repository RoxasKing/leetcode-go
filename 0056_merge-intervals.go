package leetcode

/*
  给出一个区间的集合，请合并所有重叠的区间。
*/

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		}
		return false
	})
	var index int
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= intervals[index][1] &&
			intervals[i][1] > intervals[index][1] {
			intervals[index][1] = intervals[i][1]
		} else if intervals[i][0] > intervals[index][1] {
			index++
			intervals[index] = intervals[i]
		}
	}
	return intervals[:index+1]
}
