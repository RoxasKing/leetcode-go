package main

// Difficulty:
// Easy

// Tags:
// Two Pointers

func reverseVowels(s string) string {
	isVowel := func(ch byte) bool {
		return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' ||
			ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U'
	}
	a := []byte(s)
	for l, r := 0, len(a)-1; l < r; {
		if isVowel(a[l]) && isVowel(a[r]) {
			a[l], a[r] = a[r], a[l]
			l++
			r--
		}
		for ; l < r && !isVowel(a[l]); l++ {
		}
		for ; l < r && !isVowel(a[r]); r-- {
		}
	}
	return string(a)
}
