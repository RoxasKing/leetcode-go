package main

// Difficulty:
// Easy

// Tags:
// Sliding Window

func decrypt(code []int, k int) []int {
	n := len(code)
	o := make([]int, n)
	if k == 0 {
		return o
	}
	for sum, i := 0, 0; i < n+abs(k); i++ {
		if k < 0 && i >= -k {
			o[i%n] = sum
			sum -= code[i+k]
		}
		sum += code[i%n]
		if k > 0 && i >= k {
			sum -= code[i-k]
			o[i-k] = sum
		}
	}
	return o
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
