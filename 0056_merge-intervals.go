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
		return intervals[i][0] < intervals[j][0]
	})
	var index int
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > intervals[index][1] {
			index++
			intervals[index] = intervals[i]
		} else {
			intervals[index][1] = Max(intervals[index][1], intervals[i][1])
		}
	}
	intervals = intervals[:index+1]
	return intervals
}
