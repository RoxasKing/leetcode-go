package main

func countCharacters(words []string, chars string) int {
	cnt0 := [26]int{}
	for i := range chars {
		cnt0[chars[i]-'a']++
	}
	out := 0
	for _, word := range words {
		cnt := [26]int{}
		for i := range word {
			idx := int(word[i] - 'a')
			cnt[idx]++
			if cnt[idx] > cnt0[idx] {
				goto END
			}
		}
		out += len(word)
	END:
	}
	return out
}
