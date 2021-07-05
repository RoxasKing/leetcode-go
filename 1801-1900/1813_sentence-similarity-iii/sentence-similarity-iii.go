package main

import "strings"

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	s1 := strings.Split(sentence1, " ")
	s2 := strings.Split(sentence2, " ")
	for len(s1) > 0 && len(s2) > 0 && s1[0] == s2[0] {
		s1 = s1[1:]
		s2 = s2[1:]
	}
	for len(s1) > 0 && len(s2) > 0 && s1[len(s1)-1] == s2[len(s2)-1] {
		s1 = s1[:len(s1)-1]
		s2 = s2[:len(s2)-1]
	}
	return len(s1) == 0 || len(s2) == 0
}
