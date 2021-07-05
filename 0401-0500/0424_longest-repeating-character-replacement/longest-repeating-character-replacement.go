package main

// Tags:
// Sliding Window

func characterReplacement(s string, k int) int {
	out := 0
	count := [26]int{}
	l, r := 0, 0
	for r = range s {
		count[s[r]-'A']++
		max := 0
		for i := 0; i < 26; i++ {
			max = Max(max, count[i])
		}
		if max+k < r+1-l {
			count[s[l]-'A']--
			l++
		} else {
			out = Max(out, r+1-l)
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
