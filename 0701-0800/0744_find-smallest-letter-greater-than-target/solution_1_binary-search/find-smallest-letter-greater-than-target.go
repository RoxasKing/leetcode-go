package main

import "sort"

// Difficulty:
// Easy

// Tags:
// Binary Search

func nextGreatestLetter(letters []byte, target byte) byte {
	i := sort.Search(len(letters), func(i int) bool { return letters[i] > target })
	if i == len(letters) {
		return letters[0]
	}
	return letters[i]
}
