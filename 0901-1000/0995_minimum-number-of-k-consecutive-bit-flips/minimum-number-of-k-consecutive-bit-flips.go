package main

// Tags:
// Sliding Window
func minKBitFlips(A []int, K int) int {
	out := 0
	a := make([]int, 0, len(A))
	b := make([]int, 0, len(A))

	for i := range A {
		a = append(a, A[i])
		b = append(b, A[i]^1)

		if i > K-1 {
			a = a[1:]
			b = b[1:]
		}

		if i >= K-1 && a[0] == 0 {
			a, b = b, a
			out++
		}
	}

	for i := range a {
		if a[i] == 0 {
			return -1
		}
	}

	return out
}
