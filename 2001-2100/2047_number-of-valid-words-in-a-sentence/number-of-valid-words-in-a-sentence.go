package main

// Difficulty:
// Easy

func countValidWords(sentence string) int {
	isPunctuation := func(ch byte) bool {
		return ch == '!' || ch == '.' || ch == ','
	}
	isDigit := func(ch byte) bool {
		return '0' <= ch && ch <= '9'
	}
	isLowerCaseLetter := func(ch byte) bool {
		return 'a' <= ch && ch <= 'z'
	}
	n := len(sentence)
	out := 0
	for i := 0; i < n; i++ {
		for ; i < n && sentence[i] == ' '; i++ {
		}
		if i == n {
			break
		}
		valid := true
		j := i
		for cntHyphens := 0; j < n && sentence[j] != ' '; j++ {
			if isPunctuation(sentence[j]) && j+1 < n && sentence[j+1] != ' ' || isDigit(sentence[j]) {
				valid = false
			} else if sentence[j] == '-' {
				if cntHyphens++; cntHyphens > 1 || j == i || j+1 == n || j+1 < n && !isLowerCaseLetter(sentence[j+1]) {
					valid = false
				}
			}
		}
		if valid {
			out++
		}
		i = j
	}
	return out
}
