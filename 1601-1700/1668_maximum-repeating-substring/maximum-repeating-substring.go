package main

// Difficulty:
// Easy

func maxRepeating(sequence string, word string) int {
	m, n := len(sequence), len(word)
	o := 0
	for i := 0; i+n <= m; i++ {
		c := 0
		for j := i; j+n <= m && sequence[j:j+n] == word; j += n {
			c++
		}
		if o < c {
			o = c
		}
	}
	return o
}
