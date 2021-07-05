package main

func findTheDifference(s string, t string) byte {
	count := [26]int{}
	for i := range s {
		count[s[i]-'a']++
	}
	var out byte
	for i := range t {
		count[t[i]-'a']--
		if count[t[i]-'a'] < 0 {
			out = t[i]
			break
		}
	}
	return out
}
