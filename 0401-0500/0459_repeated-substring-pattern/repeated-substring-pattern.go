package main

import "strings"

func repeatedSubstringPattern(s string) bool {
	for i := len(s) >> 1; i >= 1; i-- {
		if len(s)%i != 0 {
			continue
		}
		flag := true
		for j := 0; j <= len(s)-2*i; j += i {
			if s[j:j+i] != s[j+i:j+2*i] {
				flag = false
				break
			}
		}
		if flag {
			return true
		}
	}
	return false
}

func repeatedSubstringPattern2(s string) bool {
	tmp := s + s
	return strings.Contains(tmp[1:len(tmp)-1], s)
}
