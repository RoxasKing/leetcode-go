package main

import "strconv"

func compressString(S string) string {
	if len(S) == 0 {
		return ""
	}
	compressed := ""
	cur, cnt := "", 0
	for i := range S {
		if S[i:i+1] == cur {
			cnt++
		} else if i > 0 {
			compressed += cur + strconv.Itoa(cnt)
			cur, cnt = S[i:i+1], 1
		} else {
			cur, cnt = S[i:i+1], 1
		}
	}
	compressed += cur + strconv.Itoa(cnt)
	if len(compressed) < len(S) {
		return compressed
	}
	return S
}
