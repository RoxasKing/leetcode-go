package main

import "math/bits"

// Tags:
// Prefix Sum
// Bit Manipulation

func countPalindromicSubsequence(s string) int {
	n := len(s)
	l := make([]int, n)
	r := make([]int, n)
	lor, ror := 0, 0
	for i := 0; i < n; i++ {
		lor |= 1 << int(s[i]-'a')
		l[i] = lor
		ror |= 1 << int(s[n-1-i]-'a')
		r[n-1-i] = ror
	}

	sts := [26]int{}
	for i := 1; i < n-1; i++ {
		sts[s[i]-'a'] |= l[i-1] & r[i+1]
	}

	var out int
	for i := 0; i < 26; i++ {
		out += bits.OnesCount(uint(sts[i]))
	}
	return out
}
