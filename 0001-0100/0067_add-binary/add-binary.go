package main

func addBinary(a string, b string) string {
	arrA, arrB := []byte(a), []byte(b)
	reverse(arrA)
	reverse(arrB)
	arr := []byte{}
	remain := 0
	for len(arrA) > 0 && len(arrB) > 0 {
		remain += int(arrA[0]-'0') + int(arrB[0]-'0')
		arr = append(arr, byte(remain%2+'0'))
		remain /= 2
		arrA = arrA[1:]
		arrB = arrB[1:]
	}
	for len(arrA) > 0 {
		remain += int(arrA[0] - '0')
		arr = append(arr, byte(remain%2+'0'))
		remain /= 2
		arrA = arrA[1:]
	}
	for len(arrB) > 0 {
		remain += int(arrB[0] - '0')
		arr = append(arr, byte(remain%2+'0'))
		remain /= 2
		arrB = arrB[1:]
	}
	if remain > 0 {
		arr = append(arr, byte(remain+'0'))
	}
	reverse(arr)
	return string(arr)
}

func reverse(arr []byte) {
	l, r := 0, len(arr)-1
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}
