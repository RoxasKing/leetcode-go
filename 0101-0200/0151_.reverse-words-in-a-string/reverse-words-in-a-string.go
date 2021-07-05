package main

import "strings"

func reverseWords(s string) string {
	arr := []string{}
	n := len(s)

	for i := 0; i < n; i++ {
		if s[i] == ' ' {
			continue
		}
		j := i
		for j+1 < n && s[j+1] != ' ' {
			j++
		}
		arr = append(arr, s[i:j+1])
		i = j
	}

	for l, r := 0, len(arr)-1; l < r; l, r = l+1, r-1 {
		arr[l], arr[r] = arr[r], arr[l]
	}

	return strings.Join(arr, " ")
}
