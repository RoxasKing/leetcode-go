package main

import "unsafe"

func replaceSpaces(S string, length int) string {
	chs := make([]byte, 0, len(S))
	for i := 0; i < length; i++ {
		if S[i] == ' ' {
			chs = append(chs, '%', '2', '0')
		} else {
			chs = append(chs, S[i])
		}
	}
	return *(*string)(unsafe.Pointer(&chs))
}
