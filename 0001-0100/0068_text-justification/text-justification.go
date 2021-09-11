package main

import "strings"

func fullJustify(words []string, maxWidth int) []string {
	var out, tmp []string
	var x int
	for _, word := range words {
		if x+len(word)+len(tmp) > maxWidth {
			var s string
			if len(tmp) == 1 {
				s = tmp[0]
				for len(s) < maxWidth {
					s += " "
				}
			} else {
				spaces := maxWidth - x
				space := spaces / (len(tmp) - 1)
				remain := spaces - space*(len(tmp)-1)
				s = tmp[0]
				for _, w := range tmp[1:] {
					if remain > 0 {
						s += " "
						remain--
					}
					for i := 0; i < space; i++ {
						s += " "
					}
					s += w
				}
			}
			out = append(out, s)
			tmp = []string{}
			x = 0
		}
		tmp = append(tmp, word)
		x += len(word)
	}
	s := strings.Join(tmp, " ")
	for len(s) < maxWidth {
		s += " "
	}
	out = append(out, s)
	return out
}
