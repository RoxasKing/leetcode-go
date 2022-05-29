package main

// DifficultY:
// Medium

// Tags:
// Two Pointers

func findClosest(words []string, word1 string, word2 string) int {
	o, i, j := -1, -1, -1
	for k, w := range words {
		if w == word1 {
			i = k
		}
		if w == word2 {
			j = k
		}
		if i != -1 && j != -1 && (o == -1 || o > abs(i-j)) {
			o = abs(i - j)
		}
	}
	return o
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
