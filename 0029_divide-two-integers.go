package leetcode

/*
  给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。
  返回被除数 dividend 除以除数 divisor 得到的商。
*/

func divide(dividend int, divisor int) int {
	var (
		tmp      uint32
		udivisor uint32
		divid    uint32
		divis    uint32
	)
	if dividend > 0 {
		divid = uint32(dividend)
	} else {
		divid = uint32(-dividend)
	}
	if divisor > 0 {
		udivisor = uint32(divisor)
	} else {
		udivisor = uint32(-divisor)
	}
	divis = udivisor
	if divid < divis {
		return 0
	}
	for {
		var i uint32
		for divis < divid {
			if divis == divid || divis<<1 > divid {
				break
			}
			divis <<= 1
			i++
		}
		tmp += 1 << i
		divid = divid - divis
		divis = udivisor
		if divid < divis {
			break
		}
	}
	out := int(tmp)
	if dividend > 0 && divisor < 0 || dividend < 0 && divisor > 0 {
		out = -out
	}
	if out >= 2147483648 {
		return 2147483647
	}
	if out <= -2147483648 {
		return -2147483648
	}
	return out
}
