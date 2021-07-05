package main

func plusOne(digits []int) []int {
	reverse(digits)
	remain := 1
	for i := range digits {
		remain += digits[i]
		digits[i] = remain % 10
		remain /= 10
	}
	if remain > 0 {
		digits = append(digits, remain)
	}
	reverse(digits)
	return digits
}

func reverse(arr []int) {
	l, r := 0, len(arr)-1
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}
