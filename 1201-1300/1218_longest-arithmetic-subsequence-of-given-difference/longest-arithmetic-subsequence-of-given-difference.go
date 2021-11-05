package main

// Difficulty:
// Medium

// Tags:
// Hash

func longestSubsequence(arr []int, difference int) int {
	out := 0
	cnt := map[int]int{}
	for _, num := range arr {
		if cnt[num] = cnt[num-difference] + 1; cnt[num] > out {
			out = cnt[num]
		}
	}
	return out
}
