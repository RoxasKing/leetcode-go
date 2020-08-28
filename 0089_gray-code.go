package main

/*
  格雷编码是一个二进制数字系统，在该系统中，两个连续的数值仅有一个位数的差异。
  给定一个代表编码总位数的非负整数 n，打印其格雷编码序列。格雷编码序列必须以 0 开头。
*/

func grayCode(n int) []int {
	out := []int{0}
	head := 1
	for i := 0; i < n; i++ {
		for j := len(out) - 1; j >= 0; j-- {
			out = append(out, head+out[j])
		}
		head <<= 1
	}
	return out
}
