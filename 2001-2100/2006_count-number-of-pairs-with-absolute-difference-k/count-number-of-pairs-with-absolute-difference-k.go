package main

// Difficulty:
// Easy

// Tags:
// Hash

func countKDifference(nums []int, k int) int {
	out := 0
	freq := map[int]int{}
	for _, num := range nums {
		if cnt, ok := freq[num-k]; ok {
			out += cnt
		}
		if cnt, ok := freq[num+k]; ok {
			out += cnt
		}
		freq[num]++
	}
	return out
}
