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
		for j := 0; j < i; j++ {
			if nums[j] >= nums[i] {
				continue
			}
			if lens[j] >= lens[i] {
				lens[i] = lens[j] + 1
				cnts[i] = cnts[j]
			} else if lens[j]+1 == lens[i] {
				cnts[i] += cnts[j]
			}
		}
	}

	var max, out int
	for _, len := range lens {
		max = Max(max, len)
	}
	for i := 0; i < n; i++ {
		if lens[i] == max {
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
