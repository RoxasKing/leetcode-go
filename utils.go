package main

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Abs(n int) int {
	if n <= 0 {
		return 0 - n
	}
	return n
}

func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Pow(a, b int) int {
	res := 1
	for i := 0; i < b; i++ {
		res *= a
	}
	return res
}
