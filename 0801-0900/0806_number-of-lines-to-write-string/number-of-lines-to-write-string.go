package main

// Difficulty:
// Easy

func numberOfLines(widths []int, s string) []int {
	o := make([]int, 2)
	o[0]++
	for i := range s {
		if o[1] += widths[s[i]-'a']; o[1] > 100 {
			o[1] = widths[s[i]-'a']
			o[0]++
		}
	}
	return o
}
