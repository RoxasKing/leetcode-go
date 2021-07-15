package main

import (
	"sort"
)

// Tags:
// Binary Search

func waysToSplit(nums []int) int {
	n := len(nums)
	pSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		pSum[i] = pSum[i-1] + nums[i-1]
	}

	var out int
	for i := 1; i <= n; i++ {
		j := sort.Search(n+1-(i+1), func(j int) bool { return pSum[j+i+1]-pSum[i] >= pSum[i] }) + i + 1
		k := sort.Search(n+1-(i+1), func(k int) bool { return pSum[k+i+1]-pSum[i] > pSum[n]-pSum[k+i+1] }) + (i + 1)
		if k > n {
			k = n
		}
		if j < k {
			out += k - j
		}
	}
	return out % mod
}

var mod = int(1e9 + 7)
