package main

// Difficulty:
// Medium

// Tags:
// Hash

func numPairsDivisibleBy60(time []int) int {
	var out int
	freq := [501]int{}
	for _, t := range time {
		for sum := 60; sum < 1000; sum += 60 {
			if sum-t > 0 && sum-t <= 500 && freq[sum-t] > 0 {
				out += freq[sum-t]
			}
		}
		freq[t]++
	}
	return out
}
