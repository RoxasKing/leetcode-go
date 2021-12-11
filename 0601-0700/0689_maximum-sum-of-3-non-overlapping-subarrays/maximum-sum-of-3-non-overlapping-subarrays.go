package main

// Difficulty:
// Hard

// Tags:
// Prefix Sum
// Dynamic Programming

func maxSumOfThreeSubarrays(nums []int, k int) []int {
	n := len(nums)
	pSum := make([]int, n)
	sSum := make([]int, n)
	pSum[0], sSum[n-1] = nums[0], nums[n-1]
	for i := 1; i < n; i++ {
		pSum[i] = pSum[i-1] + nums[i]
		sSum[n-1-i] = sSum[n-i] + nums[n-1-i]
	}
	maxP := make([]int, n)
	idxP := make([]int, n)
	maxS := make([]int, n)
	idxS := make([]int, n)
	for i, maxL, idxL, maxR, idxR := 0, 0, 0, 0, n-1; i+k-1 < n; i++ {
		j := i + k - 1
		sumL, l := pSum[j], i
		sumR, r := sSum[n-1-j], n-1-j
		if i > 0 {
			sumL -= pSum[i-1]
			sumR -= sSum[n-i]
		}
		if maxL < sumL {
			maxL, idxL = sumL, l
		}
		if maxR <= sumR {
			maxR, idxR = sumR, r
		}
		maxP[j], idxP[j] = maxL, idxL
		maxS[n-1-j], idxS[n-1-j] = maxR, idxR
	}
	var sum int
	var out []int
	for i := k; i+k < n; i++ {
		cur := maxP[i-1] + pSum[i+k-1] - pSum[i-1] + maxS[i+k]
		if sum < cur {
			sum = cur
			out = []int{idxP[i-1], i, idxS[i+k]}
		}
	}
	return out
}
