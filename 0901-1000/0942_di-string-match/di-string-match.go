package main

// Difficulty:
// Easy

// Tags:
// Math

func diStringMatch(s string) []int {
	n := len(s)
	o := make([]int, 0, n+1)
	l, r := 0, n
	for _, c := range s {
		if c == 'I' {
			o = append(o, l)
			l++
		} else {
			o = append(o, r)
			r--
		}
	}
	o = append(o, l)
	return o
}
