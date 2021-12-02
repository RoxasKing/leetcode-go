package main

import "math"

// Difficulty:
// Hard

// Tags:
// Math

func poorPigs(buckets, minutesToDie, minutesToTest int) int {
	return int(math.Ceil(math.Log(float64(buckets)) / math.Log(float64(minutesToTest/minutesToDie+1))))
}
