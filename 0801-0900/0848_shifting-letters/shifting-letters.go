package main

// Tags:
// Prefix Sum

func shiftingLetters(s string, shifts []int) string {
	arr := []byte(s)
	for i, x := len(s)-1, 0; i >= 0; i-- {
		x = (x + shifts[i]) % 26
		arr[i] = (arr[i]-'a'+byte(x))%26 + 'a'
	}
	return string(arr)
}
