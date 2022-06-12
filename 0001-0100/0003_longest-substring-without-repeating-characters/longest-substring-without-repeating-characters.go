package main

// Difficulty:
// Medium

// Tags:
// Sliding Window

func lengthOfLongestSubstring(s string) int {
	h := [128]bool{}
	o := 0
	for l, r := 0, 0; r < len(s); r++ {
		for ; h[s[r]]; l++ {
			h[s[l]] = false
		}
		h[s[r]] = true
		if o < r+1-l {
			o = r + 1 - l
		}
	}
	return o
}
