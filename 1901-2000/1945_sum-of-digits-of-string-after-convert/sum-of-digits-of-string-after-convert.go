package main

func getLucky(s string, k int) int {
	arr := []int{}
	for i := range s {
		num := int(s[i] - 'a' + 1)
		for ; num > 0; num /= 10 {
			arr = append(arr, num%10)
		}
	}
	for ; k > 0; k-- {
		sum := 0
		for _, num := range arr {
			sum += num
		}
		arr = arr[:0]
		for ; sum > 0; sum /= 10 {
			arr = append(arr, sum%10)
		}
	}

	var out int
	for i := len(arr) - 1; i >= 0; i-- {
		out = out*10 + arr[i]
	}
	return out
}
