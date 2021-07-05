package main

func addToArrayForm(A []int, K int) []int {
	reverse(A)
	B := []int{}
	for K > 0 {
		B = append(B, K%10)
		K /= 10
	}
	C := make([]int, 0, Max(len(A), len(B))+1)
	remain := 0
	for len(A) > 0 || len(B) > 0 {
		if len(A) > 0 {
			remain += A[0]
			A = A[1:]
		}
		if len(B) > 0 {
			remain += B[0]
			B = B[1:]
		}
		C = append(C, remain%10)
		remain /= 10
	}
	if remain > 0 {
		C = append(C, remain)
	}
	reverse(C)
	return C
}

func reverse(arr []int) {
	l, r := 0, len(arr)-1
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
