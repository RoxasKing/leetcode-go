package main

func longestPalindrome(s string) int {
	count := [128]int{}
	for _, c := range s {
		count[c]++
	}
	out := 0
	for _, v := range count {
		out += v / 2 * 2
	}
	if out < len(s) {
		out++
	}
	return out
}
