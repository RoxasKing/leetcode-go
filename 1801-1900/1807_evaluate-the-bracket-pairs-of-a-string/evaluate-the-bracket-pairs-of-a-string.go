package main

// Tags:
// Two Pointers
func evaluate(s string, knowledge [][]string) string {
	dict := make(map[string]string)
	for _, k := range knowledge {
		dict[k[0]] = k[1]
	}
	out := ""
	l, r := 0, 0
	for r = range s {
		ch := s[r]
		if ch == '(' {
			out += s[l:r]
			l = r + 1
		} else if ch == ')' {
			if val, ok := dict[s[l:r]]; ok {
				out += val
			} else {
				out += "?"
			}
			l = r + 1
		}
	}
	out += s[l : r+1]
	return out
}
