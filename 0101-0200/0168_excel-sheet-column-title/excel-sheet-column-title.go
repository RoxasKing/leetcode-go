package main

func convertToTitle(columnNumber int) string {
	arr := []byte{}
	for ; columnNumber > 0; columnNumber /= 26 {
		i := columnNumber%26 - 1
		if i < 0 {
			i += 26
			columnNumber -= 26
		}
		arr = append(arr, byte(i)+'A')
	}
	for i := 0; i < len(arr)>>1; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	return string(arr)
}
