package main

func makeFancyString(s string) string {
	chs := make([]byte, 0, len(s))
	cnt, pre := 0, byte(' ')
	for i := range s {
		if s[i] == pre {
			cnt++
		} else {
			for cnt = Min(cnt, 2); cnt > 0; cnt-- {
				chs = append(chs, pre)
			}
			cnt, pre = 1, s[i]
		}
	}
	for cnt = Min(cnt, 2); cnt > 0; cnt-- {
		chs = append(chs, pre)
	}
	return string(chs)
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
