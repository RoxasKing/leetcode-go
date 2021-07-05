package main

import "strconv"

func findNthDigit(n int) int {
	if n < 10 {
		return n
	}
	n -= 10
	base, k := 10, 2
	for n >= 9*k*base {
		n -= 9 * k * base
		base *= 10
		k++
	}
	str := strconv.Itoa(base + n/k)
	i := n % k
	return int(str[i] - '0')
}
