package main

import (
	"fmt"
)

func myAtoi(str string) int {
	var out int
	// flg1: true 是负数 , false 是正数
	// flg2: true 已存在一个‘-’或‘+’
	// flg3: true 之前遍历的字符是数字
	var flg1, flg2, flg3 bool
	for _, s := range str {
		// 如果字符非数字或 ' ' '+' '-'
		if (s > '9' || s < '0') && s != '-' && s != '+' && s != ' ' {
			break
		}
		if s == ' ' || s == '-' || s == '+' {
			if flg2 || flg3 {
				break
			}
			if s == '-' {
				flg1 = true
				flg2 = true
			} else if s == '+' {
				flg1 = false
				flg2 = true
			}
			continue
		}
		if s >= '0' && s <= '9' {
			// 如果字符是数字
			flg3 = true
			out = out*10 + int(s-'0')
			// 如果超过 32 位有符号整数范围
			if out >= 2147483648 && flg1 {
				return -2147483648
			} else if out > 2147483647 {
				return 2147483647
			}
			continue
		}
	}
	if flg1 {
		return -out
	} else {
		return out
	}
}

func main() {
	//str := "    -88827   5655  U"
	//str := "0-1"
	//str := "test 123"
	str := "123 aaa"
	fmt.Println(myAtoi(str))
}
