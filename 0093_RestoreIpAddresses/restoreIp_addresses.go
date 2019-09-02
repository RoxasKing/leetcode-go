package main

import "fmt"

func restoreIpAddresses(s string) []string {
	if len(s) < 4 || len(s) > 12 {
		return []string{}
	}

	out := []string{}
	ip := make([]string, 4)

	var dfs func(int, int)
	dfs = func(i int, start int) {
		if i == 3 {
			tmp := s[start:]
			if isOK(tmp) {
				ip[3] = tmp
				out = append(out, IP(ip))
			}
			return
		}

		maxRemain := 3 * (3 - i)

		for end := start + 1; end <= len(s)-3+i; end++ {
			if end+maxRemain < len(s) {
				continue
			}
			if end-start > 3 {
				break
			}
			tmp := s[start:end]
			if isOK(tmp) {
				ip[i] = tmp
				dfs(i+1, end)
			}
		}
	}
	dfs(0, 0)
	return out
}

func IP(s []string) string {
	return fmt.Sprintf("%s.%s.%s.%s", s[0], s[1], s[2], s[3])
}

func isOK(s string) bool {
	if len(s) > 1 && s[0] == '0' {
		return false
	}
	if len(s) < 3 {
		return true
	}

	switch s[0] {
	case '1':
		return true
	case '2':
		if '0' <= s[1] && s[1] <= '4' {
			return true
		}
		if s[1] == '5' && '0' <= s[2] && s[2] <= '5' {
			return true
		}
	}
	return false
}

func main() {
	s := "25525511135"
	fmt.Println(restoreIpAddresses(s))
}
