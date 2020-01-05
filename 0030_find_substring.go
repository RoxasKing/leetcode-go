package My_LeetCode_In_Go

/*
  给定一个字符串 s 和一些长度相同的单词 words。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。
  注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。
*/

func findSubstring(s string, words []string) []int {
	strLen, wordsLen := len(s), len(words)
	if strLen == 0 || wordsLen == 0 {
		return nil
	}
	wordLen := len(words[0])
	if strLen < wordsLen*wordLen {
		return nil
	}
	out := make([]int, 0, strLen)
	wordsCount := make(map[string]int, wordsLen)
	for _, word := range words {
		if len(word) == wordLen {
			wordsCount[word]++
		} else {
			return nil
		}
	}
	var (
		currentCount map[string]int
		countAll     int
		left         int
		right        int
	)
	initCount := func() {
		currentCount = make(map[string]int)
		countAll = 0
	}
	rightMove := func() {
		currentCount[s[left:left+wordLen]]--
		countAll--
		left += wordLen
	}
	for i := 0; i < wordLen; i++ {
		left, right = i, i
		initCount()
		for strLen-left >= wordsLen*wordLen {
			word := s[right : right+wordLen]
			if max, ok := wordsCount[word]; !ok {
				left, right = right+wordLen, right+wordLen
				initCount()
			} else if currentCount[word]+1 > max {
				rightMove()
			} else {
				currentCount[word]++
				countAll++
				right += wordLen
				if countAll == wordsLen {
					out = append(out, left)
					rightMove()
				}
			}
		}
	}
	if len(out) == 0 {
		return nil
	}
	return out
}
