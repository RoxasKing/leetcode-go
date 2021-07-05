package main

func getSum(a int, b int) int {
	for b != 0 {
		xor := a ^ b
		and := a & b
		a = xor
		b = and << 1
	}
	return a
}
