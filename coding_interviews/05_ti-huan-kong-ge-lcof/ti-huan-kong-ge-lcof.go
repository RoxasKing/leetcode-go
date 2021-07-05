package main

func replaceSpace(s string) string {
	out := ""
	for i := range s {
		if s[i] == ' ' {
			out += "%20"
		} else {
			out += s[i : i+1]
		}
	}
	return out
}
