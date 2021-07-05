package main

// Tags:
// Bit Manipulation
func add(a int, b int) int {
	and, carry := 0, 0
	for b != 0 {
		and = a ^ b
		carry = (a & b) << 1
		a, b = and, carry
	}
	return a
}
