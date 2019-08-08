package main

import "fmt"

func merge(intervals [][]int) [][]int {
	var out [][]int
	if len(intervals) < 1 {
		return out
	}
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
	//fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	//fmt.Println(merge([][]int{{1, 4}, {4, 5}, {6, 10}, {10, 12}}))
	//fmt.Println(merge([][]int{{1, 4}, {0, 4}}))
	fmt.Println(merge([][]int{{1, 4}, {0, 0}}))
	fmt.Println(merge([][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}))
}
