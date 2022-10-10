package main

// Difficulty:
// Medium

// Tags:
// Greedy

func minAddToMakeValid(s string) int {
	o, cnt := 0, 0
	for _, c := range s {
		if c == '(' {
			cnt++
		} else {
			cnt--
		}
		if cnt < 0 {
			o++
			cnt = 0
		}
	}
	o += cnt
	return o
}
