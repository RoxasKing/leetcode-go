package main

func findSubstring(s string, words []string) []int {
	sLen, wordsLen := len(s), len(words)
	if sLen == 0 || wordsLen == 0 {
		return nil
	}
	wordLen := len(words[0])
	totalWordLen := wordLen * wordsLen
	if sLen < totalWordLen {
		return nil
	}
	wordMaxCount := make(map[string]int, wordsLen)
	for _, word := range words {
		wordMaxCount[word]++
	}
	var wordCount map[string]int
	var count, l, r int
	initCount := func() {
		wordCount = make(map[string]int)
		count = 0
	}
	shiftRight := func() {
		wordCount[s[l:l+wordLen]]--
		count--
		l += wordLen
	}
	var out []int
	for i := 0; i < wordLen; i++ {
		l, r = i, i
		initCount()
		for sLen-l >= totalWordLen {
			word := s[r : r+wordLen]
			if max, ok := wordMaxCount[word]; !ok {
				l = r + wordLen
				r += wordLen
				initCount()
			} else if wordCount[word]+1 > max {
				shiftRight()
			} else {
				wordCount[word]++
				count++
				r += wordLen
				if count == wordsLen {
					out = append(out, l)
					shiftRight()
				}
			}
		}
	}
	return out
}
