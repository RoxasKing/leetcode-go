package main

func checkOnesSegment(s string) bool {
	cnt := 0
	for s != "" {
		for s != "" && s[0] == '0' {
			s = s[1:]
		}
		count := 0
		for s != "" && s[0] == '1' {
			s = s[1:]
			count++
		}
		if count > 0 {
			cnt++
		}
		if cnt > 1 {
			return false
		}
	}
	return cnt == 1
}
