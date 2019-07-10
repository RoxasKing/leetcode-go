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
	// 获取上一个报数的值
	pre := countAndSay(n - 1)
	// 上一个报数的首位
	elem := pre[0]
	// 统计当前数字的个数,从 1 开始
	count := 1
	for i := 1; i < len(pre); i++ {
		if pre[i] == elem {
			// 遍历到的数字与当前正在统计的字符相同,则计数加 1
			count++
		} else {
			// 遍历到的数字与当前统计的字符不同，则将统计的数字的个数+数字add到输出中
			out += strconv.Itoa(count) + string(elem)
			// 将统计的数字换为当前遍历到的，并将计数重置为 1
			elem = pre[i]
			count = 1
		}
	}
	// 输出最后一位数字的个数+数字add到输出中
	out += strconv.Itoa(count) + string(elem)
	return out
}

func main() {
	in := 8
	fmt.Println(countAndSay(in))
}
