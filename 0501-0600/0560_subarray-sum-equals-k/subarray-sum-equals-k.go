package main

// Difficulty:
// Medium

// Tags:
// Prefix Sum
// Hash

func subarraySum(nums []int, k int) int {
	out, sum := 0, 0
	freq := map[int]int{0: 1}
	for _, num := range nums {
		sum += num
		if cnt, ok := freq[sum-k]; ok {
			out += cnt
		}
		freq[sum]++
	}
	return out
}
