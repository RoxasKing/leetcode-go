package main

import "fmt"

func intToRoman(num int) string {
	var runes []rune
	var i int
	for num != 0 {
		switch {
		case num >= 1000:
			i = num / 1000
			num -= i * 1000
			for ; i > 0; i-- {
				runes = append(runes, 'M')
			}
		case num >= 900:
			num -= 900
			runes = append(runes, 'C', 'M')
		case num >= 500:
			num -= 500
			runes = append(runes, 'D')
		case num >= 400:
			num -= 400
			runes = append(runes, 'C', 'D')
		case num >= 100:
			i = num / 100
			num -= i * 100
			for ; i > 0; i-- {
				runes = append(runes, 'C')
			}
		case num >= 90:
			num -= 90
			runes = append(runes, 'X', 'C')
		case num >= 50:
			num -= 50
			runes = append(runes, 'L')
		case num >= 40:
			num -= 40
			runes = append(runes, 'X', 'L')
		case num >= 10:
			i = num / 10
			num -= i * 10
			for ; i > 0; i-- {
				runes = append(runes, 'X')
			}
		case num >= 9:
			num -= 9
			runes = append(runes, 'I', 'X')
		case num >= 5:
			num -= 5
			runes = append(runes, 'V')
		case num >= 4:
			num -= 4
			runes = append(runes, 'I', 'V')
		case num >= 1:
			for ; num > 0; num-- {
				runes = append(runes, 'I')
			}
		}
	}
	return string(runes)
}
func intToRoman2(num int) string {
	out := ""
	val := [13]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	str := [13]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"};
	i := 0
	for num > 0 {
		if num < val[i] {
			i++
		} else {
			out += str[i]
			num -= val[i]
		}
	}
	return out;
}

func main() {
	i := 58
	fmt.Println(intToRoman(i))
	fmt.Println(intToRoman2(i))
}
