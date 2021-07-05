package main

// Tags:
// Sliding Window
func checkInclusion(s1 string, s2 string) bool {
	cnt := [26]int{}
	for i := range s1 {
		cnt[s1[i]-'a']++
	}
	for l, r := 0, 0; r < len(s2); r++ {
		cnt[s2[r]-'a']--
		for l <= r && cnt[s2[r]-'a'] < 0 {
			cnt[s2[l]-'a']++
			l++
		}
		if r+1-l == len(s1) {
			return true
		}
	}
	return false
}
