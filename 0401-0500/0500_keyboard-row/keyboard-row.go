package main

// Difficulty:
// Easy

// Tags:
// Hash
// Bit Manipulation

func findWords(words []string) []string {
	dict := [26]int{2, 4, 4, 2, 1, 2, 2, 2, 1, 2, 2, 2, 4, 4, 1, 1, 1, 1, 2, 1, 1, 4, 1, 4, 1, 4}
	out := make([]string, 0, len(words))
	for _, word := range words {
		t := 0
		for i := range word {
			x := word[i] - 'a'
			if 'A' <= word[i] && word[i] <= 'Z' {
				x = word[i] - 'A'
			}
			t |= dict[x]
		}
		if t == 1 || t == 2 || t == 4 {
			out = append(out, word)
		}
	}
	return out
}
