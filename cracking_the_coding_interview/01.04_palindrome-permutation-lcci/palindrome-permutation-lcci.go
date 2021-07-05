package main

// Tags:
// Hash

func canPermutePalindrome(s string) bool {
	cnt := map[rune]int{}
	for _, ch := range s {
		cnt[ch]++
	}
	count := 0
	for _, v := range cnt {
		if v&1 == 1 {
			count++
		}
		if count > 1 {
			return false
		}
	}
	return true
}
