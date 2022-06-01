package main

// Difficulty:
// Medium

// Tags:
// Bit Manipulation

func divide(dividend int, divisor int) int {
	if dividend == -1<<31 && divisor == -1 {
		return 1<<31 - 1
	}
	flg := 1
	if dividend < 0 && divisor > 0 || dividend > 0 && divisor < 0 {
		flg = -1
	}
	dividend, divisor = abs(dividend), abs(divisor)
	val := 0
	for dividend >= divisor {
		i, x := 1, divisor
		for x<<1 <= dividend {
			x <<= 1
			i <<= 1
		}
		dividend -= x
		val += i
	}
	return flg * val
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
