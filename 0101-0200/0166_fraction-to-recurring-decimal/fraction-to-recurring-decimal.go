package main

import (
	"strconv"
)

// Tags:
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
