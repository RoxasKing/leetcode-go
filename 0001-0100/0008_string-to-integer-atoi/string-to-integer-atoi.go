package main

func myAtoi(s string) int {
	i, n := 0, len(s)
	for i < n && s[i] == ' ' {
		i++
	}
	flg := 1
	if i < n && s[i] == '-' {
		flg = -1
		i++
	} else if i < n && s[i] == '+' {
		i++
	}
	sum := 0
	for i < n && '0' <= s[i] && s[i] <= '9' {
		num := int(s[i] - '0')
		if flg == 1 && ((1<<31-1-num)/10 < sum || (1<<31-1-num)/10 == sum && (1<<31-1-num)%10 == 0) {
			return 1<<31 - 1
		} else if flg == -1 && ((1<<31-num)/10 < sum || (1<<31-num)/10 == sum && (1<<31-num)%10 == 0) {
			return -1 << 31
		}
		sum = sum*10 + num
		i++
	}
	return sum * flg
}
