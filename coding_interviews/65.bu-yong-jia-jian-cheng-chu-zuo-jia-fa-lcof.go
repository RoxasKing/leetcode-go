package codinginterviews

/*
  写一个函数，求两个整数之和，要求在函数体内不得使用 “+”、“-”、“*”、“/” 四则运算符号。

  提示：
    a, b 均可能是负数或 0
    结果不会溢出 32 位整数
*/

func add(a int, b int) int {
	var sum, carry int
	for b != 0 {
		sum = a ^ b
		carry = (a & b) << 1
		a, b = sum, carry
	}
	return a
}
