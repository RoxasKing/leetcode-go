package main

import "strings"

var nameMap = map[int]string{
	1000000000: "Billion",
	1000000:    "Million",
	1000:       "Thousand",
	100:        "Hundred",
	90:         "Ninety",
	80:         "Eighty",
	70:         "Seventy",
	60:         "Sixty",
	50:         "Fifty",
	40:         "Forty",
	30:         "Thirty",
	20:         "Twenty",
	19:         "Nineteen",
	18:         "Eighteen",
	17:         "Seventeen",
	16:         "Sixteen",
	15:         "Fifteen",
	14:         "Fourteen",
	13:         "Thirteen",
	12:         "Twelve",
	11:         "Eleven",
	10:         "Ten",
	9:          "Nine",
	8:          "Eight",
	7:          "Seven",
	6:          "Six",
	5:          "Five",
	4:          "Four",
	3:          "Three",
	2:          "Two",
	1:          "One",
	0:          "Zero",
}

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	strs := []string{}
	for i := 1000000000; i >= 1; i /= 1000 {
		a := num / i
		num -= a * i
		if a == 0 {
			continue
		}

		b := a / 100
		if b > 0 {
			strs = append(strs, nameMap[b], nameMap[100])
		}

		c := a - b*100
		if c > 0 {
			if val, ok := nameMap[c]; ok { // 1~20, 30, 40, 50, 60, 70, 80, 90
				strs = append(strs, val)
			} else {
				strs = append(strs, nameMap[c/10*10], nameMap[c%10])
			}
		}

		if i > 1 {
			strs = append(strs, nameMap[i])
		}
	}
	return strings.Join(strs, " ")
}
