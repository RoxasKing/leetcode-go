package main

// Difficulty:
// Medium

// Tags:
// Backtracking

func ambiguousCoordinates(s string) []string {
	o := []string{}
	getPos := func(s string) []string {
		a := []string{}
		if s[0] != '0' || s == "0" {
			a = append(a, s)
		}
		for i := 1; i < len(s); i++ {
			if i > 1 && s[0] == '0' || s[len(s)-1] == '0' {
				continue
			}
			a = append(a, s[:i]+"."+s[i:])
		}
		return a
	}
	s = s[1 : len(s)-1]
	for i := 1; i < len(s); i++ {
		ls := getPos(s[:i])
		rs := getPos(s[i:])
		for _, l := range ls {
			for _, r := range rs {
				o = append(o, "("+l+", "+r+")")
			}
		}
	}
	return o
}
