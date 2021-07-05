package main

// Tags:
// Prefix Sum + Hash

func minSubarray(nums []int, p int) int {
	out := -1
	n := len(nums)
	sSum := make([]int, n)

	for i, sum := n-1, 0; i >= 0; i-- {
		sum += nums[i]
		sum %= p
		if sum == 0 && (out == -1 || i < out) {
			out = i
		}
		sSum[i] = sum
	}

	dict := map[int]int{0: -1}

	for i, sum := 0, 0; i < n; i++ {
		sum += nums[i]
		sum %= p
		if sum == 0 && (out == -1 || n-(i+1) < out) {
			out = n - (i + 1)
		}
		if j, ok := dict[p-sSum[i]]; ok && (out == -1 || i-j-1 < out) {
			out = i - j - 1
		}
		dict[sum] = i
	}

	return out
}
