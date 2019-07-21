package main

import "fmt"

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	if len(intervals) == 1 {
		return intervals
	}
	out := [][]int{}
	for i := 1; i < len(intervals); i++ {
		if intervals[i-1][len(intervals[i-1])-1] >= intervals[i][0] {
			tmp := []int{}
			out = append(out, append(tmp, intervals[i-1][0], intervals[i][len(intervals[i])-1]))
		} else {
			out = append(out, intervals[i-1])
		}
	}
	return out
}

func main() {
	intervals := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}
	fmt.Println(merge(intervals))
}
