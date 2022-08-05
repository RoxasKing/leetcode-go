package main

// Difficulty:
// Medium

// Tags:
// Math

func mirrorReflection(p int, q int) int {
	g := gcd(p, q)
	p, q = p/g&1, q/g&1
	if q == 0 {
		return 0
	}
	if p == 0 {
		return 2
	}
	return 1
}

func gcd(a, b int) int {
	for ; b != 0; a, b = b, a%b {
	}
	return a
}
