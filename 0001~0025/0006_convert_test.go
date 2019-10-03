package _001_0025

import (
	"fmt"
	"testing"
)

func convert(s string, numRows int) string {
	size := len(s)
	// 如果字符串长度小于等于给定行数或者给定行数为 1
	if s == "" || size <= numRows || numRows == 1 {
		return s
	}
	// 每两列字符总长度
	interval := (numRows - 1) * 2
	out := make([]byte, size)
	var k int
	// 按行号生成每行的字符
	for i := 0; i < numRows; i++ {
		// 每隔两列一循环
		for j := i; j < size; j += interval {
			out[k] = s[j]
			k++
			// 如果不是第一行或者最后一行
			if i != 0 && i != numRows-1 {
				// 计算间隔列字符位置
				next := j + interval - 2*i
				// 如果没超过字符串长度
				if next < size {
					out[k] = s[next]
					k++
				}
			}
		}
	}
	return string(out)
}

func TestConvert(t *testing.T) {
	s := "PAYPALISHIRING"
	numRow := 3
	fmt.Println(convert(s, numRow))
}
