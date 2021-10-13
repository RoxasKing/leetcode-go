package main

// Difficulty:
// Medium

// Tags:
// Bit Manipulation

func divide(dividend int, divisor int) int {
	flg := false
	if dividend < 0 && divisor > 0 || dividend > 0 && divisor < 0 {
		flg = true
	}
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}
	out := 0
	for dividend >= divisor {
		i := 1
		x := divisor
		for x<<1 <= dividend {
			x <<= 1
			i <<= 1
		}
		dividend -= x
		out += i
	}
	if flg {
		out = -out
	}
	if out > 1<<31-1 {
		return 1<<31 - 1
	}
	return out
}
