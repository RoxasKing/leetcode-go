package main

// Difficulty:
// Easy

// Tags:
// Two Pointers

func numDifferentIntegers(word string) int {
	h := map[string]struct{}{"": {}}
	o, l, r := 0, 0, 0
	for ; r < len(word); r++ {
		if ch := word[r]; '0' <= ch && ch <= '9' {
			if l < r && word[l] == '0' {
				l++
			}
		} else {
			if _, ok := h[word[l:r]]; !ok {
				h[word[l:r]] = struct{}{}
				o++
			}
			l = r + 1
		}
	}
	if _, ok := h[word[l:r]]; !ok {
		o++
	}
	return o
}
