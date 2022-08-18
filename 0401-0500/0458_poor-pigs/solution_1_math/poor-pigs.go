package main

// Difficulty:
// Hard

// Tags:
// Math

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	o := 0
	for sum, x := 1, minutesToTest/minutesToDie+1; sum < buckets; sum *= x {
		o++
	}
	return o
}
