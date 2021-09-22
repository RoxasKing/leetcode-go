package main

// Tags:
// Dynamic Programming

func findNumberOfLIS(nums []int) int {
	n := len(nums)
	cnts := make([]int, n)
	for i := range cnts {
		cnts[i] = 1
	}
	lens := make([]int, n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] >= nums[j] {
				continue
			}
			if lens[i] >= lens[j] {
				lens[j] = lens[i] + 1
				cnts[j] = cnts[i]
			} else if lens[i]+1 == lens[j] {
				cnts[j] += cnts[i]
			}
		}
	}
	var max, out int
	for _, len := range lens {
		max = Max(max, len)
	}
	for i, len := range lens {
		if len == max {
			out += cnts[i]
		}
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
