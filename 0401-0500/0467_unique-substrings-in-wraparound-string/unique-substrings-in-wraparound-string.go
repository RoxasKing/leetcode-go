package main

// Difficulty:
// Medium

// Tags:
// Dynamic Programming

func findSubstringInWraproundString(p string) int {
	f := [26]int{}
	c := 0
	for i := range p {
		if i > 0 && (int(p[i-1]-'a')+1)%26 == int(p[i]-'a') {
			c++
		} else {
			c = 1
		}
		if f[p[i]-'a'] < c {
			f[p[i]-'a'] = c
		}
	}
	o := 0
	for i := 0; i < 26; i++ {
		o += f[i]
	}
	return o
}
