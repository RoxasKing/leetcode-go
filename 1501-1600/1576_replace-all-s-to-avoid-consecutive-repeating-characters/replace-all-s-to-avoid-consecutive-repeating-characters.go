package main

// Difficulty:
// Easy

func modifyString(s string) string {
	cs := []byte(s)
	for i := range cs {
		if cs[i] != '?' {
			continue
		}
		cs[i] = 'a'
		for (i > 0 && cs[i-1] == cs[i]) || (i < len(s)-1 && cs[i] == cs[i+1]) {
			cs[i] = (cs[i]-'a'+1)%26 + 'a'
		}
	}
	return string(cs)
}
