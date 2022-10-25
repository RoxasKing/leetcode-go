package main

// Difficulty:
// Easy

// Tags:
// Two Pointers

func mergeAlternately(word1 string, word2 string) string {
	a := []byte{}
	i, j := 0, 0
	for i < len(word1) && j < len(word2) {
		if len(a)&1 == 0 {
			a = append(a, word1[i])
			i++
		} else {
			a = append(a, word2[j])
			j++
		}
	}
	return string(a) + word1[i:] + word2[j:]
}
