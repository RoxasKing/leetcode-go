package main

import "fmt"

func insert(intervals [][]int, newInterval []int) [][]int {
	out := [][]int{}
	tmp := out[0]
	for i := 0; i < len(out); i++ {

	}

	return out
}

func main() {
	intervals := [][]int{{1, 3}, {6, 9}}
	newInterval := []int{2, 5}
	fmt.Println(insert(intervals, newInterval))
}
