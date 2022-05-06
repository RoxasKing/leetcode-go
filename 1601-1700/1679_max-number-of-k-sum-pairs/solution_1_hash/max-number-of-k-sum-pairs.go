package main

// Difficulty:
// Medium

// Tags:
// Hash

func maxOperations(nums []int, k int) int {
	freq := map[int]int{}
	out := 0
	for _, x := range nums {
		if x >= k {
			continue
		}
		if v, ok := freq[k-x]; ok && v > 0 {
			freq[k-x]--
			out++
			continue
		}
		freq[x]++
	}
	return out
}
