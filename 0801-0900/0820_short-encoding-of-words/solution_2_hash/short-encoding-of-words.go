package main

// Difficulty:
// Medium

// Tags:
// Hash

func minimumLengthEncoding(words []string) int {
	mark := make(map[string]struct{})
	for _, word := range words {
		mark[word] = struct{}{}
	}
	for _, word := range words {
		for word != "" {
			word = word[1:]
			delete(mark, word)
		}
	}
	out := 0
	for w := range mark {
		out += len(w) + 1
	}
	return out
}
