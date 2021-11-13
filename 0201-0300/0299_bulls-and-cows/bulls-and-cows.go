package main

import "strconv"

// Difficulty:
// Medium

// Tags:
// Hash

func getHint(secret string, guess string) string {
	a, b := 0, 0
	freq := [10]int{}
	for i := range secret {
		if secret[i] == guess[i] {
			a++
		} else {
			freq[secret[i]-'0']++
		}
	}
	for i := range guess {
		if secret[i] == guess[i] {
			continue
		}
		if freq[guess[i]-'0'] > 0 {
			b++
			freq[guess[i]-'0']--
		}
	}
	return strconv.Itoa(a) + "A" + strconv.Itoa(b) + "B"
}
