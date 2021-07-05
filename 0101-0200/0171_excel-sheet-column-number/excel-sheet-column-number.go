package main

func titleToNumber(s string) int {
	var out int
	base := 1
	for s != "" {
		last := s[len(s)-1]
		out += base * int(last-'A'+1)
		s = s[:len(s)-1]
		base *= 26
	}
	return out
}
