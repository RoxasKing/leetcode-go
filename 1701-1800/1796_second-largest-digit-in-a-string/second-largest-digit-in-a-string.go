package main

// Difficulty:
// Easy

func secondHighest(s string) int {
	o, t := -1, -1
	for _, c := range s {
		if !('0' <= c && c <= '9') {
			continue
		}
		if t < int(c-'0') {
			o, t = t, int(c-'0')
		} else if t > int(c-'0') && o < int(c-'0') {
			o = int(c - '0')
		}
	}
	return o
}
