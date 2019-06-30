package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	out := ""
	if n == 1 {
		return "1"
	}
	pre := countAndSay(n - 1)
	elem := pre[0]
	var count int
	for i := 0; i < len(pre); i++ {
		if pre[i] == elem {
			count++
		} else {
			out += strconv.Itoa(count) + string(elem)
			elem = pre[i]
			count = 1
		}
	}
	out += strconv.Itoa(count) + string(elem)
	return out
}

func main() {
	in := 8
	fmt.Println(countAndSay(in))
}
