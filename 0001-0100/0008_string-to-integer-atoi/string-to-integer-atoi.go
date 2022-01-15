package main

// Difficulty:
// Medium

func myAtoi(s string) int {
	n := len(s)
	i := 0
	flg := 1
	for ; i < n && s[i] == ' '; i++ {
	}
	if i < n && s[i] == '+' {
		i++
	} else if i < n && s[i] == '-' {
		i++
		flg = -1
	}
	out := 0
	for ; i < n && '0' <= s[i] && s[i] <= '9'; i++ {
		x := int(s[i] - '0')
		if flg == 1 && (1<<31-1-x)/10 < out {
			return 1<<31 - 1
		} else if flg == -1 && (1<<31-x)/10 < out {
			return -1 << 31
		}
		out = out*10 + x
	}
	return flg * out
}
