package main

import "strconv"

// Difficulty:
// Medium

// Tags:
// Math
// Hash

func fractionToDecimal(numerator int, denominator int) string {
	out := ""
	if numerator != 0 && numerator^denominator < 0 {
		out += "-"
	}
	if numerator < 0 {
		numerator = -numerator
	}
	if denominator < 0 {
		denominator = -denominator
	}
	div := func() {
		out += strconv.Itoa(numerator / denominator)
		numerator %= denominator
	}
	if div(); numerator == 0 {
		return out
	}
	out += "."
	dict := map[int]int{}
	for numerator > 0 {
		if i, ok := dict[numerator]; ok {
			return out[:i] + "(" + out[i:] + ")"
		}
		dict[numerator] = len(out)
		numerator *= 10
		div()
	}
	return out
}
