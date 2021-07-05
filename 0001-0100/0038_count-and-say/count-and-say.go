package main

import (
	"strconv"
)

func countAndSay(n int) string {
	out := "1"
	for i := 1; i < n; i++ {
		newOut, cur, count := "", out[0:1], 1
		for j := 1; j < len(out); j++ {
			if out[j:j+1] == cur {
				count++
				continue
			}
			newOut += strconv.Itoa(count) + cur
			cur, count = out[j:j+1], 1
		}
		newOut += strconv.Itoa(count) + cur
		out = newOut
	}
	return out
}
