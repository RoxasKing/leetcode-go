package main

func makeEqual(words []string) bool {
	n := len(words)
	cnt := [26]int{}
	for _, word := range words {
		for i := range word {
			cnt[word[i]-'a']++
		}
	}
	for i := 0; i < 26; i++ {
		if cnt[i]%n > 0 {
			return false
		}
	}
	return true
}
