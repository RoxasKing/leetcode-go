package main

// Difficulty:
// Easy

func reversePrefix(word string, ch byte) string {
	chs := []byte(word)
	for i := range chs {
		if chs[i] == ch {
			for j := 0; j <= i>>1; j++ {
				chs[j], chs[i-j] = chs[i-j], chs[j]
			}
			break
		}
	}
	return string(chs)
}
