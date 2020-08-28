package main

/*
  给定一个字符串 s 和一些长度相同的单词 words。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。
  注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。
*/

func findSubstring(s string, words []string) []int {
	if len(s) == 0 || len(words) == 0 {
		return nil
	}
	if len(s) < len(words)*len(words[0]) {
		return nil
	}
	out := make([]int, 0, len(s))
	maxCount := make(map[string]int, len(words))
	for _, word := range words {
		if len(word) == len(words[0]) {
			maxCount[word]++
		} else {
			return nil
		}
	}
	var (
		curCount map[string]int
		countAll int
		left     int
		right    int
	)
	initCount := func() {
		curCount = make(map[string]int)
		countAll = 0
	}
	shiftRight := func() {
		curCount[s[left:left+len(words[0])]]--
		countAll--
		left += len(words[0])
	}
	for i := 0; i < len(words[0]); i++ {
		left, right = i, i
		initCount()
		for len(s)-left >= len(words)*len(words[0]) {
			word := s[right : right+len(words[0])]
			if max, ok := maxCount[word]; !ok {
				left, right = right+len(words[0]), right+len(words[0])
				initCount()
			} else if curCount[word]+1 > max {
				shiftRight()
			} else {
				curCount[word]++
				countAll++
				right += len(words[0])
				if countAll == len(words) {
					out = append(out, left)
					shiftRight()
				}
			}
		}
	}
	if len(out) == 0 {
		return nil
	}
	return out
}
