package leetcode

/*
  给出一个无重叠的 ，按照区间起始端点排序的区间列表。
  在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。
*/

func insert(intervals [][]int, newInterval []int) [][]int {
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
	out := make([][]int, len(intervals)+1+left-right)
	copy(out[:left], intervals[:left])
	out[left] = newInterval
	copy(out[left+1:], intervals[right:])
	return out
}

func insert2(intervals [][]int, newInterval []int) [][]int {
	var l, r int
	for r < len(intervals) {
		if intervals[r][0] > newInterval[1] {
			break
		}
		if intervals[r][1] < newInterval[0] {
			l++
		} else {
			if intervals[r][0] < newInterval[0] {
				newInterval[0] = intervals[r][0]
			}
			if intervals[r][1] > newInterval[1] {
				newInterval[1] = intervals[r][1]
			}
		}
		r++
	}
	out := make([][]int, len(intervals)-(r-l-1))
	copy(out[:l], intervals[:l])
	out[l] = newInterval
	copy(out[l+1:], intervals[r:])
	return out
}
