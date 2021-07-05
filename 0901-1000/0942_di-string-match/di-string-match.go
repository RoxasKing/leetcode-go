package main

// Tags:
// Math
func diStringMatch(s string) []int {
	n := len(s)
	out := make([]int, 0, n+1)
	l, r := 0, n
	for i := range s {
		if s[i] == 'I' {
			out = append(out, l)
			l++
		} else {
			out = append(out, r)
			r--
		}
	}
	out = append(out, l)
	return out
}
