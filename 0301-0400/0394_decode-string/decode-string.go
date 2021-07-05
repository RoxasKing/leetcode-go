package main

import (
	"strconv"
	"strings"
)

// Tags:
// Stack
func decodeString(s string) string {
	stack := []string{}
	for r := 0; r < len(s); {
		if '0' <= s[r] && s[r] <= '9' {
			l := r
			for '0' <= s[r] && s[r] <= '9' {
				r++
			}
			stack = append(stack, s[l:r])
		} else if s[r] == ']' {
			tmp := []string{}
			for stack[len(stack)-1] != "[" {
				tmp = append(tmp, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			for i := 0; i < len(tmp)/2; i++ {
				tmp[i], tmp[len(tmp)-i-1] = tmp[len(tmp)-i-1], tmp[i]
			}
			stack = stack[:len(stack)-1] // delete "["
			times, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-1]
			stack = append(stack, strings.Repeat(strings.Join(tmp, ""), times))
			r++
		} else {
			stack = append(stack, s[r:r+1])
			r++
		}
	}
	return strings.Join(stack, "")
}
