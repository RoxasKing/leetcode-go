package main

// Difficulty:
// Easy

func longestNiceSubstring(s string) string {
	var out string
	n := len(s)
	for l := 0; l < n; l++ {
		for r := l; r < n; r++ {
			has1 := [26]bool{}
			has2 := [26]bool{}
			for i := l; i <= r; i++ {
				if 'a' <= s[i] && s[i] <= 'z' {
					has1[s[i]-'a'] = true
				} else {
					has2[s[i]-'A'] = true
				}
			}
			ok := true
			for i := 0; i < 26; i++ {
				if has1[i] && !has2[i] || !has1[i] && has2[i] {
					ok = false
					break
				}
			}
			if ok && len(out) < r+1-l {
				out = s[l : r+1]
			}
		}
	}
	return out
}
