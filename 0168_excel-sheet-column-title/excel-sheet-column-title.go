package main

/*
  给定一个正整数，返回它在 Excel 表中相对应的列名称。
*/

func convertToTitle(n int) string {
	var out string
	for n != 0 {
		remain := n % 26
		if remain == 0 {
			out = "Z" + out
			n--
		} else {
			out = string('A'+byte(remain-1)) + out
		}
		n /= 26
	}
	return out
}
