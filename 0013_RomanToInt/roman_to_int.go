package main

import "fmt"

func romanToInt(s string) int {
	m := map[string]int{
		"I":  1,
		"V":  5,
		"IV": 4,
		"X":  10,
		"IX": 9,
		"L":  50,
		"XL": 40,
		"C":  100,
		"XC": 90,
		"D":  500,
		"CD": 400,
		"M":  1000,
		"CM": 900,
	}
	out := 0
	for i := len(s) - 1; i >= 0; i-- {
		if i-1 >= 0 && m[string(s[i-1])] < m[string(s[i])] {
			out += m[string(s[i-1:i+1])]
			i--
		} else {
			out += m[string(s[i])]
		}
	}
	return out
}

func main() {
	str := "III"
	fmt.Println(romanToInt(str))
}
