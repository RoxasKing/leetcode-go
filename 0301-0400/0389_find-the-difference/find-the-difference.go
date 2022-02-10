package main

// Difficulty:
// Easy

func findTheDifference(s string, t string) byte {
	freq := [26]int{}
	for i := range s {
		freq[s[i]-'a']++
	}
	for i := range t {
		if freq[t[i]-'a'] == 0 {
			return t[i]
		}
		freq[t[i]-'a']--
	}
	return ' '
}
