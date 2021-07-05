package main

func lengthOfLastWord(s string) int {
	var count int
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			count++
		} else if count != 0 {
			break
		}
	}
	return count
}
