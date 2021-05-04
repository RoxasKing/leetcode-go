package main

import (
	"strconv"
)

/*
  Given two integers representing the numerator and denominator of a fraction, return the fraction in string format.

  If the fractional part is repeating, enclose the repeating part in parentheses.

  If multiple answers are possible, return any of them.

  It is guaranteed that the length of the answer string is less than 104 for all the given inputs.

  Example 1:
    Input: numerator = 1, denominator = 2
    Output: "0.5"

  Example 2:
    Input: numerator = 2, denominator = 1
    Output: "2"

  Example 3:
    Input: numerator = 2, denominator = 3
    Output: "0.(6)"

  Example 4:
    Input: numerator = 4, denominator = 333
    Output: "0.(012)"

  Example 5:
    Input: numerator = 1, denominator = 5
    Output: "0.2"

  Constraints:
    1. -2^31 <= numerator, denominator <= 2^31 - 1
    2. denominator != 0

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/fraction-to-recurring-decimal
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Math + Hash
func fractionToDecimal(numerator int, denominator int) string {
	out := ""
	if numerator < 0 && denominator > 0 || numerator > 0 && denominator < 0 {
		out += "-"
	}
	numerator = Abs(numerator)
	denominator = Abs(denominator)
	out += strconv.Itoa(numerator / denominator)
	remain := numerator % denominator
	if remain == 0 {
		return out
	}

	out += "."
	memo := make(map[int]int)
	for remain > 0 {
		if i, ok := memo[remain]; ok {
			return out[:i] + "(" + out[i:] + ")"
		}
		memo[remain] = len(out)
		remain *= 10
		out += strconv.Itoa(remain / denominator)
		remain %= denominator
	}
	return out
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
