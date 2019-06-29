package main

import "fmt"

func myAtoi(str string) int {
	if str == "" {
		return 0
	}
	out := 0
	var flg1, flg2, flg3 bool
	for _, s := range str {
		if s == ' ' {
			if flg3 || flg2 {
				break
			}
			continue
		}
		if s > '9' && s < '0' && s != '-' {
			return 0
		}
		if s == '-' {
			if flg3 {
				break
			}
			if flg2 {
				return 0
			}
			if flg3 {
				return 0
			}
			flg1 = true
			flg2 = true
			continue
		}
		if s == '+' {
			if flg3 {
				break
			}
			if flg2 == true {
				return 0
			}
			if flg3 {
				return 0
			}
			flg1 = false
			flg2 = true
			continue
		}
		if s >= '0' && s <= '9' {
			flg3 = true
			out = out*10 + int(s-'0')
			if out >= 2147483648 && flg1 {
				return -2147483648
			}
			if out > 2147483647 {
				return 2147483647
			}
			continue
		}
		if s > '9' || s < '0' {
			break
		}
	}
	if flg1 {
		return -out
	}
	if out > 0 {
		return out
	}
	return 0
}

func main() {
	str := "    -88827   5655  U"
	fmt.Println(myAtoi(str))
}
