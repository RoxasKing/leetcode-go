package main

// Difficulty:
// Easy

// Tags:
// Hash

func canConstruct(ransomNote string, magazine string) bool {
	freq := [26]int{}
	for i := range magazine {
		freq[magazine[i]-'a']++
	}
	for i := range ransomNote {
		if freq[ransomNote[i]-'a']--; freq[ransomNote[i]-'a'] < 0 {
			return false
		}
	}
	return true
}
