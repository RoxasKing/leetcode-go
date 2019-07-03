package main

import (
	"fmt"
)

func letterCombinations(digits string) []string {
	size := len(digits)
	if size < 1 {
		return nil
	}
	var out []string
	var tmp string
	dict := map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}
	// 存放每个字符串游标记录
	save := make(map[int]int, size)
	for i := 0; i < size; i++ {
		save[i] = 0
	}
	point := 0
	for {
		// 通过每个字符游标记录，组合成新的字符串
		for i := 0; i < size; i++ {
			tmp += string(dict[digits[i]][save[i]])
		}
		out = append(out, tmp)
		tmp = ""
		point = 0
		// 循环判断每个字符的游标是否到顶
		for size-1-point >= 0 {
			save[size-1-point]++
			// 如果游标没有到顶
			if save[size-1-point] < len(dict[digits[size-1-point]]) {
				// 继续组合成新的字符串
				break
			} else {
				// 如果第一个字符的游标没有到顶
				if save[0] < len(dict[digits[0]]) {
					// 将最后一个字符的游标重置
					save[size-1-point] = 0
					// 继续判断前一个字符的游标记录
					point++
					continue
				} else {
					return out
				}
			}
		}
	}
}

func main() {
	digits := "234"
	fmt.Println(letterCombinations(digits))
}
