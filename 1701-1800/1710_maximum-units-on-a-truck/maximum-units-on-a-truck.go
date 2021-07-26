package main

import (
	"sort"
)

// Tags:
// Greedy

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool { return boxTypes[i][1] > boxTypes[j][1] })
	var out int
	for truckSize > 0 && len(boxTypes) > 0 {
		if truckSize >= boxTypes[0][0] {
			out += boxTypes[0][0] * boxTypes[0][1]
			truckSize -= boxTypes[0][0]
			boxTypes = boxTypes[1:]
		} else {
			out += truckSize * boxTypes[0][1]
			truckSize = 0
		}
	}
	return out
}
