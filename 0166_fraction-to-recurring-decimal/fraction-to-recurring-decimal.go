package main

import (
	"fmt"
	"strconv"
)

/*
  给定两个整数，分别表示分数的分子 numerator 和分母 denominator，以字符串形式返回小数。

  如果小数部分为循环小数，则将循环的部分括在括号内。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/fraction-to-recurring-decimal
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}
	var out string
	if numerator < 0 && denominator > 0 || numerator > 0 && denominator < 0 {
		out += "-"
	}
	dividend := Abs(numerator)
	divisor := Abs(denominator)
	out += strconv.Itoa(dividend / divisor)
	remain := dividend % divisor
	if remain == 0 {
		return out
	}
	out += "."
	dict := make(map[int]int)
	repeatPos := -1
	for remain != 0 {
		if pos, ok := dict[remain]; ok {
			repeatPos = pos
			break
		} else {
			dict[remain] = len(out)
		}
		remain *= 10
		out += strconv.Itoa(remain / divisor)
		remain %= divisor
	}
	if repeatPos == -1 {
		return out
	}
	return fmt.Sprintf("%s(%s)", out[0:repeatPos], out[repeatPos:])
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
