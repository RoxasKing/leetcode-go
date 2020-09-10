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
	var isNeg bool
	if dividend > 0 && divisor < 0 || dividend < 0 && divisor > 0 {
		isNeg = true
	}
	udividend := uint32(Abs(dividend))
	udivisor := uint32(Abs(divisor))
	var out int
	tmpDivisor := udivisor
	for udividend >= udivisor {
		var n int
		for tmpDivisor < udividend {
			if tmpDivisor<<1 > udividend {
				break
			}
			tmpDivisor <<= 1
			n++
		}
		udividend -= tmpDivisor
		tmpDivisor = udivisor
		cur := 1 << n
		tmp := 1<<31 - 1 - cur
		if out >= tmp+1 && isNeg {
			return -1 << 31
		} else if out >= tmp && !isNeg {
			return 1<<31 - 1
		}
		out += cur
	}
	if isNeg {
		return -out
	}
	return out
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
