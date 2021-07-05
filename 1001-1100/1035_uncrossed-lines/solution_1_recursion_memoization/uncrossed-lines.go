package main

// Tags:
// Recursion
// Memoization

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	memo := map[int]int{}
	n1, n2 := len(nums1), len(nums2)
	return dp(memo, nums1, nums2, n1, n2, 0, 0)
}

func dp(memo map[int]int, nums1, nums2 []int, n1, n2, i, j int) int {
	if i == n1 || j == n2 {
		return 0
	}

	if val, ok := memo[i*500+j]; ok {
		return val
	}

	out := 0

	nj := j
	for nj < n2 && nums1[i] != nums2[nj] {
		nj++
	}
	if nj < n2 {
		out = Max(out, 1+dp(memo, nums1, nums2, n1, n2, i+1, nj+1))
	}

	ni := i
	for ni < n1 && nums1[ni] != nums2[j] {
		ni++
	}
	if ni < n1 {
		out = Max(out, 1+dp(memo, nums1, nums2, n1, n2, ni+1, j+1))
	}

	out = Max(out, dp(memo, nums1, nums2, n1, n2, i+1, j+1))
	memo[i*500+j] = out

	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
