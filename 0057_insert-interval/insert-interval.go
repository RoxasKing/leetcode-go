package main

import "sort"

/*
  给出一个无重叠的 ，按照区间起始端点排序的区间列表。
  在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。

  示例 1：
    输入：intervals = [[1,3],[6,9]], newInterval = [2,5]
    输出：[[1,5],[6,9]]

  示例 2：
    输入：intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
    输出：[[1,2],[3,10],[12,16]]
    解释：这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10] 重叠。

    注意：输入类型已在 2019 年 4 月 15 日更改。请重置为默认代码定义以获取新的方法签名。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/insert-interval
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func insert(intervals [][]int, newInterval []int) [][]int {
	n := len(intervals)
	index := sort.Search(n, func(i int) bool { return intervals[i][0] >= newInterval[0] })
	out := make([][]int, 0, n+1)
	if index == 0 {
		out = append(out, newInterval)
	} else {
		out = append(out, intervals[:index]...)
		index--
		intervals[index] = newInterval
	}
	for i := index; i < n; i++ {
		if intervals[i][0] > out[len(out)-1][1] {
			out = append(out, intervals[i])
		} else if intervals[i][1] > out[len(out)-1][1] {
			out[len(out)-1][1] = intervals[i][1]
		}
	}
	return out
}
