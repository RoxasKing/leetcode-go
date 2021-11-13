package main

// Difficulty:
// Easy

func findPoisonedDuration(timeSeries []int, duration int) int {
	out := 0
	last := 0
	for _, t := range timeSeries {
		out += duration
		if t < last {
			out += t - last
		}
		last = t + duration
	}
	return out
}
