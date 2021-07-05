package main

// Tags:
// Hash

func findMaxLength(nums []int) int {
	out := 0
	dict := map[int]int{0: -1}
	diff := 0
	for i, num := range nums {
		if num == 0 {
			diff--
		} else {
			diff++
		}
		if j, ok := dict[diff]; ok {
			out = Max(out, i-j)
		} else {
			dict[diff] = i
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
