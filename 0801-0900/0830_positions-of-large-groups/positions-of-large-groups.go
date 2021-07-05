package main

func largeGroupPositions(s string) [][]int {
	out := [][]int{}
	l, r := 0, 0
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			r++
		} else {
			if r+1-l >= 3 {
				out = append(out, []int{l, r})
			}
			l, r = i, i
		}
	}
	if r+1-l >= 3 {
		out = append(out, []int{l, r})
	}
	return out
}
