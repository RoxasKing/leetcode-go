package main

import (
	"strings"
)

// Difficulty:
// Easy

func reorderSpaces(text string) string {
	n, cnt := len(text), 0
	words := []string{}
	for l, r := 0, 0; l < n; l = r {
		for ; l < n && text[l] == ' '; l++ {
			cnt++
		}
		for r = l; r < n && text[r] != ' '; r++ {
		}
		if l == r {
			break
		}
		words = append(words, text[l:r])
	}
	if len(words) == 1 {
		return words[0] + strings.Repeat(" ", cnt)
	}
	spaces := strings.Repeat(" ", cnt/(len(words)-1))
	o := strings.Join(words, spaces)
	if cnt%(len(words)-1) != 0 {
		o += strings.Repeat(" ", cnt%(len(words)-1))
	}
	return o
}
