package main

import (
	"fmt"
)

func isMatch(s string, p string) bool {
	slen := len(s)
	plen := len(p)

	// sp[i][j] 代表了 s[:i] 能否与 p[:j]
	sp := make([][]bool, slen+1)
	for i := range sp {
		sp[i] = make([]bool, plen+1)
	}

	// sp[0][0] 代表 s = "" 且 p = "" 时
	sp[0][0] = true

	// 根据题目的设定， "" 可以与 "a*b*c*" 相匹配,所以，需要把相应的 sp 设置成 true
	for j := 1; j < plen && sp[0][j-1]; j += 2 {
		if p[j] == '*' {
			sp[0][j+1] = true
		}
	}

	for i := 0; i < slen; i++ {
		for j := 0; j < plen; j++ {
			if p[j] == '.' || p[j] == s[i] {
				// p[j] 与 s[i] 可以匹配上，所以，只要前面匹配，这里就能匹配上
				sp[i+1][j+1] = sp[i][j]
			} else if p[j] == '*' {
				/* 此时，p[j] 的匹配情况与 p[j-1] 的内容相关。 */
				if p[j-1] != s[i] && p[j-1] != '.' {
					/**
					 * p[j] 无法与 s[i] 匹配上
					 * p[j-1:j+1] 只能被当做 ""
					 */
					sp[i+1][j+1] = sp[i+1][j-1]
				} else {
					/**
					 * p[j] 与 s[i] 匹配上
					 * p[j-1;j+1] 作为 "x*", 可以有三种解释
					 */
					sp[i+1][j+1] = sp[i+1][j-1] || /* "x*" 解释为 "" */
						sp[i+1][j] || /* "x*" 解释为 "x" */
						sp[i][j+1] /* "x*" 解释为 "xx..." */
				}
			}
		}
	}
	return sp[slen][plen]
}

func main() {
	var s, p string
	s = "aab"
	p = "c*a*b"
	fmt.Println(isMatch(s, p))
}
