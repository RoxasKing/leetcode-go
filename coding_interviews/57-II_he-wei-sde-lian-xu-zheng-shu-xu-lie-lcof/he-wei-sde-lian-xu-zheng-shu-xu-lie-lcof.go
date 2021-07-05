package main

func findContinuousSequence(target int) [][]int {
	out := [][]int{}
	l, r := 1, 2
	for l < r {
		sum := (l + r) * (r + 1 - l) / 2
		if sum < target {
			r++
		} else if sum > target {
			l++
		} else {
			arr := make([]int, 0, r+1-l)
			for i := l; i <= r; i++ {
				arr = append(arr, i)
			}
			out = append(out, arr)
			l++
		}
	}
	return out
}
