package main

import (
	"math"
	"strconv"
	"strings"
)

// Tags:
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
