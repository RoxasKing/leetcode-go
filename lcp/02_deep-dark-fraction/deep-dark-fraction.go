package main

func fraction(cont []int) []int {
	out := []int{1, 0}
	for len(cont) > 0 {
		top := len(cont) - 1
		num := cont[top]
		cont = cont[:top]
		out[0], out[1] = out[1]+out[0]*num, out[0]
	}
	return out
}
