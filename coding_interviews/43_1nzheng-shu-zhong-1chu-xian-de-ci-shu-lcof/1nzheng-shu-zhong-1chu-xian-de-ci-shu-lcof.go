package main

// Tags:
// Math
func countDigitOne(n int) int {
	out := 0
	for i := 1; i <= n; i *= 10 {
		div := i * 10
		out += (n / div) * i // e.g. if i == 10, (1234/100)*10 = 12*10, left side 0~11, right side 0~9
		remain := n % div
		if remain >= i<<1-1 { // e.g. if i == 10, 1234%100 = 34 >= 19, so 1210 ~ 1219 have i number
			out += i
		} else if remain >= i { // e.g. if i == 10, 1215%100 = 15 >= 10, so 1210 ~ 1215 have 15%i+1 = 6 number
			out += remain%i + 1
		}
	}
	return out
}
