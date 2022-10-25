package main

// Difficulty:
// Easy

// Tags:
// Hash

func checkIfPangram(sentence string) bool {
	h := [26]int{}
	cnt := 0
	for _, c := range sentence {
		if h[c-'a'] == 0 {
			cnt++
		}
		h[c-'a']++
	}
	return cnt == 26
}
