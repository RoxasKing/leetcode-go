package main

func mergeAlternately(word1 string, word2 string) string {
	i, j := 0, 0
	out := ""
	flg := 0
	for i < len(word1) && j < len(word2) {
		if flg == 0 {
			out += word1[i : i+1]
			i++
		} else {
			out += word2[j : j+1]
			j++
		}
		flg ^= 1
	}
	out += word1[i:]
	out += word2[j:]
	return out
}
