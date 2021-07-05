package main

// Tags:
// Sliding Window
func numOfSubarrays(arr []int, k int, threshold int) int {
	target := k * threshold
	out := 0
	sum := 0
	for i := range arr {
		sum += arr[i]
		if i > k-1 {
			sum -= arr[i-k]
		}
		if i >= k-1 && sum >= target {
			out++
		}
	}
	return out
}
