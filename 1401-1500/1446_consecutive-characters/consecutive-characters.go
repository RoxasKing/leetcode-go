package main

// Difficulty:
// Easy

func maxPower(s string) int {
	out, cnt := 1, 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			cnt++
			if cnt > out {
				out = cnt
			}
		} else {
			cnt = 1
		}
	}
	return out
}
