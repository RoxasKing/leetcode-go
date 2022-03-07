package main

// Difficulty:
// Easy

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	flg := false
	if num < 0 {
		flg = true
		num = -num
	}
	arr := []byte{}
	for ; num > 0; num /= 7 {
		arr = append(arr, byte(num%7)+'0')
	}
	if flg {
		arr = append(arr, '-')
	}
	for i := 0; i < len(arr)>>1; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	return string(arr)
}
