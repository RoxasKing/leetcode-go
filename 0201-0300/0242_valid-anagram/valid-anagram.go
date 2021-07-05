package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := [26]int{}
	for i := range s {
		count[s[i]-'a']++
	}
	for i := range t {
		count[t[i]-'a']--
		if count[t[i]-'a'] < 0 {
			return false
		}
	}
	return true
}
