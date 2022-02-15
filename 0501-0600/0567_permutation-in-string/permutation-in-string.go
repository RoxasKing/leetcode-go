package main

// Difficulty:
// Medium

// Tags:
// Sliding Window

func checkInclusion(s1 string, s2 string) bool {
	freq := [26]int{}
	for i := range s1 {
		freq[s1[i]-'a']++
	}
	for l, r := 0, 0; r < len(s2); r++ {
		for freq[s2[r]-'a']--; freq[s2[r]-'a'] < 0; l++ {
			freq[s2[l]-'a']++
		}
		if r+1-l == len(s1) {
			return true
		}
	}
	return false
}
