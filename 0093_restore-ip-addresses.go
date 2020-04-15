package leetcode

import (
	"strings"
)

/*
  给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。
*/

// DFS
func restoreIPAddresses(s string) []string {
	if len(s) < 4 || 12 < len(s) {
		return nil
	}
	isOK := func(str string) bool {
		switch {
		case len(str) > 1 && str[0] == '0':
			return false
		case len(str) < 3:
			return true
		case len(str) == 3:
			if str[0] == '1' {
				return true
			} else if str[0] == '2' {
				if '0' <= str[1] && str[1] <= '4' {
					return true
				} else if str[1] == '5' && '0' <= str[2] && str[2] <= '5' {
					return true
				}
			}
		}
		return false
	}
	var out []string
	var dfs func([]string, int)
	dfs = func(ip []string, l int) {
		if len(ip) == 3 {
			if isOK(s[l:]) {
				ip = append(ip, s[l:])
				out = append(out, strings.Join(ip, "."))
			}
			return
		}
		maxRemain := 3 * (3 - len(ip))
		for r := l + 1; r <= l+3 && r <= len(s)-3+len(ip); r++ {
			if r+maxRemain < len(s) {
				continue
			}
			if isOK(s[l:r]) {
				ip = append(ip, s[l:r])
				dfs(ip, r)
				ip = ip[:len(ip)-1]
			}
		}
	}
	dfs(make([]string, 0, 4), 0)
	return out
}
