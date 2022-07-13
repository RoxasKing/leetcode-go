package main

import "strings"

// Difficulty:
// Medium

// Tags:
// Hash

func replaceWords(dict []string, sentence string) string {
	hasPrefix := make(map[string]bool, len(dict))
	hasLength := make([]bool, 101)
	for i := range dict {
		hasPrefix[dict[i]] = true
		hasLength[len(dict[i])] = true
	}
	lengths := make([]int, 101)
	var index int
	for i, ok := range hasLength {
		if ok {
			lengths[index] = i
			index++
		}
	}
	lengths = lengths[:index]

	words := strings.Split(sentence, " ")
	for i, word := range words {
		for j := 0; j < len(lengths) && lengths[j] < len(word); j++ {
			if hasPrefix[word[:lengths[j]]] {
				words[i] = word[:lengths[j]]
				break
			}
		}
	}
	return strings.Join(words, " ")
}
