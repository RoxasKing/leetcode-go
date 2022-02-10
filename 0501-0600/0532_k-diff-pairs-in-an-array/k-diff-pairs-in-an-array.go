package main

// Difficulty:
// Medium

// Tags:
// Hash

func findPairs(nums []int, k int) int {
	out := 0
	freq := map[int]int{}
	vis := map[int]struct{}{}
	for _, num := range nums {
		freq[num]++
		if k == 0 {
			if freq[num] == 2 {
				out++
			}
			continue
		}
		if _, ok := vis[num]; ok {
			continue
		}
		if _, ok := freq[num-k]; ok {
			vis[num] = struct{}{}
			vis[num-k] = struct{}{}
			out++
		}
		if _, ok := freq[num+k]; ok {
			vis[num] = struct{}{}
			vis[num+k] = struct{}{}
			out++
		}
	}
	return out
}
