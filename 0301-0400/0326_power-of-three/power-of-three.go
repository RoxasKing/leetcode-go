package main

import (
	"strconv"
	"strings"
)

/*
  给定一个整数，写一个函数来判断它是否是 3 的幂次方。
*/

func isPowerOfThree(n int) bool {
	if n < 1 {
		return false
	}
	s := strconv.FormatInt(int64(n), 3)
	return s[0] == '1' && strings.Count(s, "0") == len(s)-1
}
