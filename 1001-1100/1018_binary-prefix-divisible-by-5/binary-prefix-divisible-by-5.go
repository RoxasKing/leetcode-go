package main

func prefixesDivBy5(A []int) []bool {
	n := len(A)
	out := make([]bool, n)
	num := 0
	for i, a := range A {
		num = (num<<1 + a) % 5
		if num == 0 {
			out[i] = true
		}
	}
	return out
}
