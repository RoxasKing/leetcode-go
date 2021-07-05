package main

func detectCapitalUse(word string) bool {
	f := word[0]
	word = word[1:]
	if isUpperCase(f) {
		cnt := 0
		for i := range word {
			if isUpperCase(word[i]) {
				cnt++
			}
		}
		if 0 < cnt && cnt < len(word) {
			return false
		}
	} else {
		for i := range word {
			if isUpperCase(word[i]) {
				return false
			}
		}
	}
	return true
}

func isUpperCase(ch byte) bool {
	return 'A' <= ch && ch <= 'Z'
}
