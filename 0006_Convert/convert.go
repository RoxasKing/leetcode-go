package main

import "fmt"

func convert(s string, numRows int) string {
	size := len(s)
	if s == "" || size <= numRows || numRows == 1 {
		return s
	}
	interval := (numRows - 1) * 2
	out := make([]byte, size)
	var k int
	for i := 0; i < numRows; i++ {
		for j := i; j < size; j += interval {
			out[k] = s[j]
			k++
			if i != 0 && i != numRows-1 {
				next := j + interval - 2*i
				if next < size {
					out[k] = s[next]
					k++
				}
			}
		}
	}
	return string(out)
}

func main() {
	s := "PAYPALISHIRING"
	numRow := 3
	fmt.Println(convert(s, numRow))
}
