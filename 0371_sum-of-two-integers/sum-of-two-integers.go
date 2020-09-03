package main

/*
  不使用运算符 + 和 - ​​​​​​​，计算两整数 ​​​​​​​a 、b ​​​​​​​之和。
*/

func getSum(a int, b int) int {
	for b != 0 {
		xor := a ^ b
		and := a & b
		a = xor
		b = and << 1
	}
	return a
}
