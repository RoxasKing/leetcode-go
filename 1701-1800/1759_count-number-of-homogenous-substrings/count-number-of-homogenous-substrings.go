package main

func countHomogenous(s string) int {
	out, cnt := 0, 1
	n := len(s)
	for i := 1; i < n; i++ {
		if s[i] == s[i-1] {
			cnt++
		} else {
			out += cnt * (cnt + 1) / 2
			out %= 1e9 + 7
			cnt = 1
		}
	}
	out += cnt * (cnt + 1) / 2
	out %= 1e9 + 7
	return out
}
