package main

// Tags:
// Sliding Window
func longestSubstring(s string, k int) int {
	out := 0
	for t := 1; t <= 26; t++ { // limit on the number of character types
		cnt := [26]int{}
		typ := 0 // current number of character types
		lss := 0 // frequency > 0 && < k 's character
		for l, r := 0, 0; r < len(s); r++ {
			c := s[r] - 'a'
			if cnt[c] == 0 {
				typ++
				lss++
			}
			cnt[c]++
			if cnt[c] == k {
				lss--
			}
			for typ > t {
				c := s[l] - 'a'
				if cnt[c] == k {
					lss++
				}
				cnt[c]--
				if cnt[c] == 0 {
					typ--
					lss--
				}
				l++
			}
			if lss == 0 {
				out = Max(out, r+1-l)
			}
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
