package main

// Difficulty:
// Easy

// Tags:
// Hash

func uniqueOccurrences(arr []int) bool {
	valCount := make([]int, 2001)
	for _, a := range arr {
		valCount[a+1000]++
	}
	countMark := make([]int, 1001)
	for _, count := range valCount {
		if count == 0 {
			continue
		}
		if countMark[count] != 0 {
			return false
		}
		countMark[count] = 1
	}
	return true
}
