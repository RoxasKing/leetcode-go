package main

func grayCode(n int) []int {
	out := []int{0}
	head := 1
	for i := 0; i < n; i++ {
		for j := len(out) - 1; j >= 0; j-- {
			out = append(out, head+out[j])
		}
		head <<= 1
	}
	return out
}
