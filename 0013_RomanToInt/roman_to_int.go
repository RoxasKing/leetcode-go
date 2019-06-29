package main

import "fmt"

func romanToInt(s string) int {
	str := [13]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"};
	val := [13]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	i, j, out := 0, 0, 0
	for {
		if i+1 < len(s) && s[i:i+2] == str[j] {
			out += val[j]
			i += 2
		} else if i < len(s) && s[i:i+1] == str[j] {
			out += val[j]
			i++
		} else {
			j++
		}
		if i >= len(s) {
			break
		}
	}
	return out
}
func romanToInt2(s string) int {
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
	return out;
}

func main() {
	str := "III"
	fmt.Println(romanToInt(str))
}
