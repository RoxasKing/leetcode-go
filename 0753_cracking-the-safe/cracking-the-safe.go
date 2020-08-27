package main

import (
	"math"
	"strconv"
	"strings"
)

/*
  有一个需要密码才能打开的保险箱。密码是 n 位数, 密码的每一位是 k 位序列 0, 1, ..., k-1 中的一个 。
  你可以随意输入密码，保险箱会自动记住最后 n 位输入，如果匹配，则能够打开保险箱。
  举个例子，假设密码是 "345"，你可以输入 "012345" 来打开它，只是你输入了 6 个字符.
  请返回一个能打开保险箱的最短字符串。

  提示：
    n 的范围是 [1, 4]。
    k 的范围是 [1, 10]。
    k^n 最大可能为 4096。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/cracking-the-safe
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// DFS + Eulerian path + Hash + Graph
func crackSafe(n int, k int) string {
	if n == 1 {
		var str string
		for i := 0; i < k; i++ {
			str += strconv.Itoa(i)
		}
		return str
	}
	passwords := make([]string, k)
	for i := 0; i < k; i++ {
		passwords[i] = strconv.Itoa(i)
	}
	for i := 1; i < n; i++ {
		tmp := make([]string, 0, len(passwords)*k)
		for _, password := range passwords {
			for j := 0; j < k; j++ {
				tmp = append(tmp, password+strconv.Itoa(j))
			}
		}
		passwords = tmp
	}
	dict := make(map[string][]string)
	for _, p := range passwords {
		k := p[:n-1]
		e := p[n-1:]
		dict[k] = append(dict[k], e)
	}
	var last string
	bytes := make([]byte, 0, len(passwords)+1)
	var dfs func(string)
	dfs = func(prefix string) {
		for len(dict[prefix]) != 0 {
			c := dict[prefix][0]
			dict[prefix] = dict[prefix][1:]
			dfs(prefix[1:] + c)
		}
		last = prefix
		bytes = append(bytes, prefix[0])
	}
	dfs(passwords[0][:n-1])
	for i := 0; i < len(bytes)>>1; i++ {
		bytes[i], bytes[len(bytes)-1-i] = bytes[len(bytes)-1-i], bytes[i]
	}
	return string(bytes) + last[1:]
}

// Hash
func crackSafe2(n int, k int) string {
	out := strings.Repeat("0", n)
	dict := map[string]bool{out: true}
	times := int(math.Pow(float64(k), float64(n)))
	for i := 0; i < times; i++ {
		prefix := out[len(out)-n+1:]
		for j := k - 1; j >= 0; j-- {
			lastNum := strconv.Itoa(j)
			password := prefix + lastNum
			if !dict[password] {
				dict[password] = true
				out += lastNum
				break
			}
		}
	}
	return out
}
