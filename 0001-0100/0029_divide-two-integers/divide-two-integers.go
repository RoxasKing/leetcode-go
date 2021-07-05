package main

func divide(dividend int, divisor int) int {
	flag := 1
	if dividend > 0 && divisor < 0 || dividend < 0 && divisor > 0 {
		flag = -1
	}
	udividend, udivisor := uint32(Abs(dividend)), uint32(Abs(divisor))
	var out int
	for div := udivisor; udividend >= udivisor; div = udivisor {
		var n int
		for div < udividend {
			if div<<1 > udividend {
				break
			}
			div <<= 1
			n++
		}
		udividend -= div
		subtracted := 1 << n
		maxRemain := 1<<31 - 1 - subtracted
		if flag == 1 && out >= maxRemain {
			return 1<<31 - 1
		} else if flag == -1 && out > maxRemain {
			return -1 << 31
		}
		out += subtracted
	}
	return flag * out
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
