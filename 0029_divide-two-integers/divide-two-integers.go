package main

/*
  给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。
  返回被除数 dividend 除以除数 divisor 得到的商。
  整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2

  提示：
    被除数和除数均为 32 位有符号整数。
    除数不为 0。
    假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−231,  231 − 1]。本题中，如果除法结果溢出，则返回 231 − 1。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/divide-two-integers
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func divide(dividend int, divisor int) int {
	var (
		quotient  uint32
		udivisor  uint32
		udividend uint32
		divis     uint32
	)
	if dividend > 0 {
		udividend = uint32(dividend)
	} else {
		udividend = uint32(-dividend)
	}
	if divisor > 0 {
		udivisor = uint32(divisor)
	} else {
		udivisor = uint32(-divisor)
	}
	divis = udivisor
	if udividend < divis {
		return 0
	}
	for {
		var i uint32
		for divis < udividend {
			if divis<<1 > udividend {
				break
			}
			divis <<= 1
			i++
		}
		quotient += 1 << i
		udividend = udividend - divis
		divis = udivisor
		if udividend < divis {
			break
		}
	}
	if dividend > 0 && divisor < 0 || dividend < 0 && divisor > 0 {
		if quotient > 1<<31 {
			return -1 << 31
		} else {
			return -int(quotient)
		}
	}
	if quotient > 1<<31-1 {
		return 1<<31 - 1
	}
	return int(quotient)
}
