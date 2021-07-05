package main

import "strings"

func sortSentence(s string) string {
	strs := strings.Split(s, " ")
	out := make([]string, len(strs))
	for _, str := range strs {
		idx := int(str[len(str)-1]-'0') - 1
		out[idx] = str[:len(str)-1]
	}
	return strings.Join(out, " ")
}
