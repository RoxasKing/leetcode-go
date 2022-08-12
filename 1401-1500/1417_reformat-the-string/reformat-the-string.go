package main

// Difficulty:
// Easy

func reformat(s string) string {
	a, b := []rune{}, []rune{}
	for _, c := range s {
		if '0' <= c && c <= '9' {
			a = append(a, c)
		} else {
			b = append(b, c)
		}
	}
	if abs(len(a)-len(b)) > 1 {
		return ""
	}
	if len(a) < len(b) {
		a, b = b, a
	}
	o := make([]rune, len(s))
	for i := 0; i < len(s); i, a = i+2, a[1:] {
		o[i] = a[0]
	}
	for i := 1; i < len(s); i, b = i+2, b[1:] {
		o[i] = b[0]
	}
	return string(o)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
