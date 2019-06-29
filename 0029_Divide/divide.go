package main

import "fmt"

func divide(dividend int, divisor int) int {
	var tmp uint32
	var udivisor uint32
	var divid uint32
	var divis uint32
	if dividend > 0 {
		divid = uint32(dividend)
	} else {
		divid = uint32(-dividend)
	}
	if divisor > 0 {
		udivisor = uint32(divisor)
	} else {
		udivisor = uint32(-divisor)
	}
	divis = udivisor

	if divid < divis {
		return 0
	}

	for {
		var i uint32
		for divis < divid {
			if divis == divid || divis<<1 > divid {
				break
			}
			divis <<= 1
			i++
		}
		tmp += 1 << i
		divid = divid - divis
		divis = udivisor
		if divid < divis {
			break
		}
	}

	out := int(tmp)
	if dividend > 0 && divisor < 0 || dividend < 0 && divisor > 0 {
		out = -out
	}
	if out >= 2147483648 {
		return 2147483647
	}
	if out <= -2147483648 {
		return -2147483648
	}

	return out
}

func main() {
	dividend := -2147483648
	divisor := -3
	fmt.Println(divide(dividend, divisor))
	//fmt.Println(1 << 6)
}
