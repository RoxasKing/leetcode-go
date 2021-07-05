package main

// Tags:
// Sliding Window
func findShortestSubArray(nums []int) int {
	n := len(nums)

	max := 0
	cnt := map[int]int{}
	for _, num := range nums {
		cnt[num]++
		max = Max(max, cnt[num])
	}

	out := n
	cnt = map[int]int{}
	for l, r := 0, 0; r < n; r++ {
		cnt[nums[r]]++
		if cnt[nums[r]] < max {
			continue
		}
		for l <= r && nums[l] != nums[r] {
			cnt[nums[l]]--
			l++
		}
		out = Min(out, r+1-l)
	}

	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
