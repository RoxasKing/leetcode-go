package main

// Difficulty:
// Easy

// Tags:
// Stack

func backspaceCompare(s string, t string) bool {
	a, b := []byte{}, []byte{}
	for i := range s {
		if s[i] != '#' {
			a = append(a, s[i])
			continue
		}
		if len(a) > 0 {
			a = a[:len(a)-1]
		}
	}
	for i := range t {
		if t[i] != '#' {
			b = append(b, t[i])
			continue
		}
		if len(b) > 0 {
			b = b[:len(b)-1]
		}
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
