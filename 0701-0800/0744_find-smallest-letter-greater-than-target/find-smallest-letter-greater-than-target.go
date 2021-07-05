package main

// Tags:
// Binary Search
func nextGreatestLetter(letters []byte, target byte) byte {
	n := len(letters)
	l, r := 0, n-1
	for l < r {
		m := (l + r) >> 1
		if letters[m] <= target {
			l = m + 1
		} else {
			r = m
		}
	}

	if letters[l] <= target {
		return letters[0]
	}
	return letters[l]
}
