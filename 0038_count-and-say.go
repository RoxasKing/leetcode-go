package My_LeetCode_In_Go

import (
	"strconv"
)

/*
  「外观数列」是一个整数序列，从数字 1 开始，序列中的每一项都是对前一项的描述。前五项如下：
       1.     1
       2.     11
       3.     21
       4.     1211
       5.     111221
    1 被读作  "one 1"  ("一个一") , 即 11。
    11 被读作 "two 1s" ("两个一"）, 即 21。
    21 被读作 "one 2",  "one 1" （"一个二" ,  "一个一") , 即 1211。
  给定一个正整数 n（1 ≤ n ≤ 30），输出外观数列的第 n 项。
  注意：整数序列中的每一项将表示为一个字符串。
*/

func countAndSay(n int) string {
	var (
		out   string
		tmp   string
		count int
		pre   string
	)
	for i := 0; i < n; i++ {
		if i == 0 {
			out = "1"
			continue
		}
		tmp, count, pre = "", 1, out[0:1]
		for j := 1; j < len(out); j++ {
			if out[j:j+1] == pre {
				count++
				continue
			}
			tmp += strconv.Itoa(count) + pre
			pre = out[j : j+1]
			count = 1
		}
		out = tmp + strconv.Itoa(count) + pre
	}
	return out
}
