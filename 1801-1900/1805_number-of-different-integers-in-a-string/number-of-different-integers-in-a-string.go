package main

// Tags:
// Two Pointers
func numDifferentIntegers(word string) int {
	mark := map[string]bool{"": true}
	cnt := 0
	l, r := 0, 0
	for ; r < len(word); r++ {
		if '0' <= word[r] && word[r] <= '9' {
			if l < r && word[l] == '0' {
				l++
			}
		} else {
			if !mark[word[l:r]] {
				cnt++
				mark[word[l:r]] = true
			}
			l = r + 1
		}
	}
	if !mark[word[l:r]] {
		cnt++
	}
	return cnt
}
