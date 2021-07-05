package main

// Tags:
// Bit Manipulation

func totalHammingDistance(nums []int) int {
	out := 0
	cnt := [32][2]int{}
	for _, num := range nums {
		for i := 31; i >= 0; i-- {
			idx := (num >> i) & 1
			out += cnt[i][1-idx]
			cnt[i][idx]++
		}
	}
	return out
}
