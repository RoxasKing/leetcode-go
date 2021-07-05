package main

// Tags:
// Two Pointers + Sliding Window
func smallestRange(nums [][]int) []int {
	mark := make(map[int][]int)
	min, max := 1<<31-1, -1<<31
	for i, arr := range nums {
		for _, num := range arr {
			mark[num] = append(mark[num], i)
			min = Min(min, num)
			max = Max(max, num)
		}
	}

	out := []int{min, max}
	cnt := make([]int, len(nums)) // 统计每个列表被包含的数
	inside := 0                   // 统计包含区间内数的列表个数
	for l, r := min, min; r <= max; r++ {
		if len(mark[r]) == 0 {
			continue
		}
		for _, idx := range mark[r] {
			if cnt[idx] == 0 {
				inside++
			}
			cnt[idx]++
		}
		for inside == len(nums) {
			if r-l < out[1]-out[0] {
				out = []int{l, r}
			}
			for _, idx := range mark[l] {
				cnt[idx]--
				if cnt[idx] == 0 {
					inside--
				}
			}
			l++
		}
	}
	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
