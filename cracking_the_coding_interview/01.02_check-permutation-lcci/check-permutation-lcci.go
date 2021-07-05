package main

import "sort"

func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	chs1 := []rune(s1)
	chs2 := []rune(s2)
	sort.Slice(chs1, func(i, j int) bool { return chs1[i] < chs1[j] })
	sort.Slice(chs2, func(i, j int) bool { return chs2[i] < chs2[j] })
	for i := 0; i < len(chs1); i++ {
		if chs1[i] != chs2[i] {
			return false
		}
	}
	return true
}
