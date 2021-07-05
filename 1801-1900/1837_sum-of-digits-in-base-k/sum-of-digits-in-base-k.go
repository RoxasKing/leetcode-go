package main

func sumBase(n int, k int) int {
	out := 0
	for n > 0 {
		out += n % k
		n /= k
	}
	return out
}
