package main

// Tags:
// Sliding Window
func findAnagrams(s string, p string) []int {
	cnt := [26]int{}
	for _, c := range p {
		cnt[c-'a']++
	}
	out := []int{}
	for l, r := 0, 0; r < len(s); r++ {
		cnt[s[r]-'a']--
		for l <= r && cnt[s[r]-'a'] < 0 {
			cnt[s[l]-'a']++
			l++
		}
		if r+1-l == len(p) {
			out = append(out, l)
			cnt[s[l]-'a']++
			l++
		}
	}
	return out
}
