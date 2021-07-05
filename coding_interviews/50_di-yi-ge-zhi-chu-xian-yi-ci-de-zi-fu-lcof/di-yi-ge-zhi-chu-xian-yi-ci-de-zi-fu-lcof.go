package main

func firstUniqChar(s string) byte {
	cnt := [128]int{}
	for i := range s {
		cnt[s[i]]++
	}
	for i := range s {
		if cnt[s[i]] == 1 {
			return s[i]
		}
	}
	return ' '
}
