package main

/*
  颠倒给定的 32 位无符号整数的二进制位。

  提示：
    1. 请注意，在某些语言（如 Java）中，没有无符号整数类型。在这种情况下，输入和输出都将被指定为有符号整数类型，并且不应影响您的实现，因为无论整数是有符号的还是无符号的，其内部的二进制表示形式都是相同的。
    2. 在 Java 中，编译器使用二进制补码记法来表示有符号整数。因此，在上面的 示例 2 中，输入表示有符号整数 -3，输出表示有符号整数 -1073741825。

  进阶:
  如果多次调用这个函数，你将如何优化你的算法？

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/reverse-bits
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Bit Operation
func reverseBits(num uint32) uint32 {
	var out uint32
	for i := 0; i < 31; i++ {
		if num&1 == 1 {
			out++
		}
		num >>= 1
		out <<= 1
	}
	if num == 1 {
		out++
	}
	return out
}
