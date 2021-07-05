package main

// Tags:
// Math

func countDigitOne(n int) int {
	out := 0
	for base := 1; n >= base; base *= 10 {
		l := n / (base * 10)
		m := n % (base * 10) / base
		r := n % base
		out += l * base
		if m == 1 {
			out += r + 1
		} else if m > 1 {
			out += base
		}
	}
	return out
}
