package main

import "strconv"

func getHint(secret string, guess string) string {
	a, b := 0, 0
	h := [10]int{}
	for i := range secret {
		h[secret[i]-'0']++
	}
	for i := range guess {
		if guess[i] == secret[i] {
			a++
			h[secret[i]-'0']--
		}
	}
	for i := range guess {
		if guess[i] == secret[i] {
			continue
		}
		if h[guess[i]-'0'] > 0 {
			b++
			h[guess[i]-'0']--
		}
	}
	return strconv.Itoa(a) + "A" + strconv.Itoa(b) + "B"
}
