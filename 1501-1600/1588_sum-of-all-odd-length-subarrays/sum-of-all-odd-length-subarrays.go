package main

// Tags:
// Prefix Sum

func sumOddLengthSubarrays(arr []int) int {
	n := len(arr)
	pSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		pSum[i+1] = pSum[i] + arr[i]
	}

	var out int
	for l := 1; l <= n; l += 2 {
		for i := 0; i+l <= n; i++ {
			out += pSum[i+l] - pSum[i]
		}
	}
	return out
}
