package main

// Difficulty:
// Easy

func maxLengthBetweenEqualCharacters(s string) int {
	o := -1
	for ch := 'a'; ch <= 'z'; ch++ {
		l, r := -1, -1
		for i, c := range s {
			if c != ch {
				continue
			}
			if l == -1 {
				l = i
			}
			r = i
		}
		o = max(o, r-l-1)
	}
	return o
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
