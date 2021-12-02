package main

// Difficulty:
// Medium

// Tags:
// Sliding Window

func findAnagrams(s string, p string) []int {
	freq := [26]int{}
	for i := range p {
		freq[p[i]-'a']++
	}
	out := []int{}
	for l, r := 0, 0; r < len(s); r++ {
		freq[s[r]-'a']--
		for ; l <= r && freq[s[r]-'a'] < 0; l++ {
			freq[s[l]-'a']++
		}
		if r+1-l == len(p) {
			out = append(out, l)
		}
	}
	return out
}
