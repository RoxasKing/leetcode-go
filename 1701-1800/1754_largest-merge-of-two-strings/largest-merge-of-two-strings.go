package main

func largestMerge(word1 string, word2 string) string {
	m, n := len(word1), len(word2)
	arr := make([]byte, 0, m+n)
	for word1 != "" && word2 != "" {
		if word1[0] > word2[0] || word1[0] == word2[0] && word1 > word2 {
			arr = append(arr, word1[0])
			word1 = word1[1:]
		} else {
			arr = append(arr, word2[0])
			word2 = word2[1:]
		}
	}
	return string(arr) + word1 + word2
}
