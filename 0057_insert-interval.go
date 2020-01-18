package My_LeetCode_In_Go

/*
  给出一个无重叠的 ，按照区间起始端点排序的区间列表。
  在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。
*/

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	out := make([][]int, 0, len(intervals)+1)
	var left, right int
	for ; right < len(intervals); right++ {
		if intervals[right][0] > newInterval[1] {
			break
		} else if intervals[right][1] < newInterval[0] {
			left++
		} else {
			if intervals[right][0] < newInterval[0] {
				newInterval[0] = intervals[right][0]
			}
			if intervals[right][1] > newInterval[1] {
				newInterval[1] = intervals[right][1]
			}
		}
	}
	out = append(out, intervals[:left]...)
	out = append(out, newInterval)
	out = append(out, intervals[right:]...)
	return out
}
