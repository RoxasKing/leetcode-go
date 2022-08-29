package main

// Difficulty:
// Easy

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	o := 0
	for i := range startTime {
		if startTime[i] <= queryTime && queryTime <= endTime[i] {
			o++
		}
	}
	return o
}
