package main

// Difficulty:
// Easy

// Tags:
// Hash

func uniqueMorseRepresentations(words []string) int {
	strs := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	set := map[string]struct{}{}
	for _, word := range words {
		str := ""
		for _, c := range word {
			str += strs[c-'a']
		}
		set[str] = struct{}{}
	}
	return len(set)
}
