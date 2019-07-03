package main

import "fmt"

func multiply(num1 string, num2 string) string {
	var out []byte
	len1 := len(num1)
	len2 := len(num2)
	next := 0
	for i := len1 - 1; i >= 0; i-- {
		for j := len2 - 1; j >= 0; j-- {
			tmp := int(num1[i]-'0')*int(num2[j]-'0') + next
			out = append(out, byte((tmp%10)+'0'))
			next = tmp / 10
		}
	}
	for i := 0; i < len(out); i++ {
		out[i], out[len(out)-i-1] = out[len(out)-i-1], out[i]
	}
	return string(out)
}

func main() {
	num1 := "12"
	num2 := "12"
	fmt.Println(multiply(num1, num2))
}
