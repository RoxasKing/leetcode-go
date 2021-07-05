package main

func reverseVowels(s string) string {
	chs := []byte(s)
	n := len(chs)
	l, r := 0, n-1
	for l < r {
		if isVowel(chs[l]) && isVowel(chs[r]) {
			chs[l], chs[r] = chs[r], chs[l]
			l++
			r--
		}
		for l < r && !isVowel(chs[l]) {
			l++
		}
		for l < r && !isVowel(chs[r]) {
			r--
		}
	}
	return string(chs)
}

func isVowel(ch byte) bool {
	return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' ||
		ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U'
}
