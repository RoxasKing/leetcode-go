package main

/*
  给出一个区间的集合，请合并所有重叠的区间。

  示例 1:
    输入: intervals = [[1,3],[2,6],[8,10],[15,18]]
    输出: [[1,6],[8,10],[15,18]]
    解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

  示例 2:
    输入: intervals = [[1,4],[4,5]]
    输出: [[1,5]]
    解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。

  注意：输入类型已于2019年4月15日更改。 请重置默认代码定义以获取新方法签名。

  提示：
    intervals[i][0] <= intervals[i][1]


  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/merge-intervals
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 || len(intervals[0]) == 0 {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	index := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > intervals[index][1] {
			index++
			intervals[index] = intervals[i]
		} else if intervals[i][1] > intervals[index][1] {
			intervals[index][1] = intervals[i][1]
		}
	}
	return intervals[:index+1]
}
