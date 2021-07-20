package main

import "strings"

func canBeTypedWords(text string, brokenLetters string) int {
	broke := [26]bool{}
	for i := range brokenLetters {
		broke[brokenLetters[i]-'a'] = true
	}

	var out int
	strs := strings.Split(text, " ")
	for _, str := range strs {
		ok := true
		for i := range str {
			if broke[str[i]-'a'] {
				ok = false
				break
			}
		}
		if ok {
			out++
		}
	}
	return out
}
