package main

import "sort"

// Difficulty:
// Easy

// Tags:
// Sorting
// Greedy

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool { return boxTypes[i][1] > boxTypes[j][1] })
	o := 0
	for ; truckSize > 0 && len(boxTypes) > 0; boxTypes = boxTypes[1:] {
		c := boxTypes[0][0]
		if c > truckSize {
			c = truckSize
		}
		truckSize -= c
		o += c * boxTypes[0][1]
	}
	return o
}
