package main

func longestNiceSubstring(s string) string {
	out := ""
	n := len(s)
	for i := 0; i < n; i++ {
		f1 := [26]bool{}
		f2 := [26]bool{}
		for j := i; j < n; j++ {
			if 'a' <= s[j] && s[j] <= 'z' {
				f1[s[j]-'a'] = true
			} else {
				f2[s[j]-'A'] = true
			}

			ok := true
			for i := 0; i < 26; i++ {
				if f1[i] && !f2[i] || !f1[i] && f2[i] {
					ok = false
					break
				}
			}

			if ok && j+1-i > len(out) {
				out = s[i : j+1]
			}
		}
	}
	return out
}
