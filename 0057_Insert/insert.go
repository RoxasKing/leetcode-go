package main

import "fmt"

func insert2(intervals [][]int, newInterval []int) [][]int {
	var left, right [][]int
	for i := 0; i < len(intervals); i++ {
		if intervals[i][1] < newInterval[0] {
			left = append(left, intervals[i])
		} else if intervals[i][0] > newInterval[1] {
			right = append(right, intervals[i])
		} else {
			if intervals[i][0] < newInterval[0] {
				newInterval[0] = intervals[i][0]
			}
			if intervals[i][1] > newInterval[1] {
				newInterval[1] = intervals[i][1]
			}
		}
	}
	return append(append(left, newInterval), right...)
}

func insert(intervals [][]int, newInterval []int) [][]int {
	var out [][]int
	intervals = append(intervals, newInterval)
	sort(intervals)
	tmp := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if tmp[1] < intervals[i][0] {
			// 如果 tmp 的区间和 intervals[i] 没有相交，
			// 将 tmp 加入到 out 中，
			// 并将 tmp 替换为 intervals[i]
			out = append(out, tmp)
			tmp = intervals[i]
		} else if tmp[1] < intervals[i][1] {
			// 如果相交，则合并区间
			tmp[1] = intervals[i][1]
		}
	}
	out = append(out, tmp)
	return out
}

func sort(intervals [][]int) {
	if len(intervals) <= 1 {
		return
	}
	mid := intervals[0][0]
	head, tail := 0, len(intervals)-1
	for i := 1; i <= tail; {
		if intervals[i][0] > mid {
			intervals[i], intervals[tail] = intervals[tail], intervals[i]
			tail--
		} else {
			intervals[i], intervals[head] = intervals[head], intervals[i]
			head++
			i++
		}
	}
	sort(intervals[:head])
	sort(intervals[head+1:])
}

func main() {
	//intervals := [][]int{{1, 3}, {6, 9}}
	//newInterval := []int{2, 5}
	//intervals := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
	//newInterval := []int{4, 8}
	intervals := [][]int{{1, 5}}
	newInterval := []int{2, 7}
	fmt.Println(insert(intervals, newInterval))
}
