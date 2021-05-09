package main

import (
	"strings"
)

/*
  Given a string s containing only digits, return all possible valid IP addresses that can be obtained from s. You can return them in any order.

  A valid IP address consists of exactly four integers, each integer is between 0 and 255, separated by single dots and cannot have leading zeros. For example, "0.1.2.201" and "192.168.1.1" are valid IP addresses and "0.011.255.245", "192.168.1.312" and "192.168@1.1" are invalid IP addresses.

  Example 1:
    Input: s = "25525511135"
    Output: ["255.255.11.135","255.255.111.35"]

  Example 2:
    Input: s = "0000"
    Output: ["0.0.0.0"]

  Example 3:
    Input: s = "1111"
    Output: ["1.1.1.1"]

  Example 4:
    Input: s = "010010"
    Output: ["0.10.0.10","0.100.1.0"]

  Example 5:
    Input: s = "101023"
    Output: ["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]

  Constraints:
    1. 0 <= s.length <= 3000
    2. s consists of digits only.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/restore-ip-addresses
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Backtracking
func restoreIpAddresses(s string) []string {
	out := []string{}
	dfs(s, len(s), 0, []string{}, &out)
	return out
}

func dfs(s string, n, i int, cur []string, out *[]string) {
	if i == n {
		if len(cur) != 4 {
			return
		}
		*out = append(*out, strings.Join(cur, "."))
		return
	}

	if (4-len(cur))*3 < n-i {
		return
	}

	cur = append(cur, s[i:i+1])
	dfs(s, n, i+1, cur, out)
	cur = cur[:len(cur)-1]

	if s[i] == '0' || i+1 == n {
		return
	}

	cur = append(cur, s[i:i+2])
	dfs(s, n, i+2, cur, out)
	cur = cur[:len(cur)-1]

	if i+2 == n {
		return
	}

	if s[i] > '2' || s[i] == '2' && (s[i+1] > '5' || s[i+1] == '5' && s[i+2] > '5') {
		return
	}

	cur = append(cur, s[i:i+3])
	dfs(s, n, i+3, cur, out)
	// cur = cur[:len(cur)-1]
}
