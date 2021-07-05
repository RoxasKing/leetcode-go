package main

func reverseWords(s string) string {
	bytes := []byte(s)
	var l, r int
	for l < len(s) {
		r = l + 1
		for r < len(s) && bytes[r] != ' ' {
			r++
		}
		reverseLetters(bytes[l:r])
		l = r + 1
	}
	return string(bytes)
}

func reverseLetters(letters []byte) {
	for i := 0; i < len(letters)>>1; i++ {
		letters[i], letters[len(letters)-1-i] = letters[len(letters)-1-i], letters[i]
	}
}
