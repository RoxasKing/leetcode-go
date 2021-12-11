package main

// Difficulty:
// Easy

// Tags:
// Hash

func shortestCompletingWord(licensePlate string, words []string) string {
	freq := [26]int{}
	for i := range licensePlate {
		letter := licensePlate[i]
		if letter == ' ' || '0' <= letter && letter <= '9' {
			continue
		}
		if 'A' <= letter && letter <= 'Z' {
			letter += 'a' - 'A'
		}
		freq[letter-'a']++
	}
	var out string
	for _, word := range words {
		cnt := [26]int{}
		for i := range word {
			cnt[word[i]-'a']++
		}
		ok := true
		for i := 0; i < 26; i++ {
			if freq[i] > cnt[i] {
				ok = false
				break
			}
		}
		if ok && (out == "" || len(word) < len(out)) {
			out = word
		}
	}
	return out
}
