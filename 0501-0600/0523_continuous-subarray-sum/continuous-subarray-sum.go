package main

// Difficulty:
// Medium

// Tags:
// Hash
// Prefix Sum

func checkSubarraySum(nums []int, k int) bool {
	h := map[int]int{0: -1}
	sum := 0
	for j, x := range nums {
		sum = (sum + x) % k
		if i, ok := h[sum]; !ok {
			h[sum] = j
		} else if j-i > 1 {
			return true
		}
	}
	return false
}
