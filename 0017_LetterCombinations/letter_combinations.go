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
		for i := 0; i < size; i++ {
			tmp += string(dict[digits[i]][save[i]])
		}
		out = append(out, tmp)
		tmp = ""
		point = 0
		for size-1-point >= 0 {
			save[size-1-point]++
			if save[size-1-point] < len(dict[digits[size-1-point]]) {
				break
			} else {
				if save[0] < len(dict[digits[0]]) {
					save[size-1-point] = 0
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
