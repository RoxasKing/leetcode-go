package main

/*
  给定一个字符串 s 和一些长度相同的单词 words。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。
  注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。

  示例 1：
    输入：
      s = "barfoothefoobarman",
      words = ["foo","bar"]
    输出：[0,9]
    解释：
    从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
    输出的顺序不重要, [9,0] 也是有效答案。

  示例 2：
    输入：
      s = "wordgoodgoodgoodbestword",
      words = ["word","good","best","word"]
    输出：[]

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

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
