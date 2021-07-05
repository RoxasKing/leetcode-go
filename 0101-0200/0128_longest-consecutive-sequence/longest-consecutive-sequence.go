package main

// Tags:
// Hash
func longestConsecutive(nums []int) int {
	contain := make(map[int]bool)
	for _, num := range nums {
		contain[num] = true
	}
	out := 0
	for _, num := range nums {
		if contain[num-1] {
			continue
		}
		cnt := 1
		for i := num + 1; contain[i]; i++ {
			cnt++
		}
		out = Max(out, cnt)
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
