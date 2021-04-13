package main

/*
  输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。

  示例 1:
    输入: n = 1
    输出: [1,2,3,4,5,6,7,8,9]

  说明：
    1. 用返回一个整数列表来代替打印
    2. n 为正整数

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/da-yin-cong-1dao-zui-da-de-nwei-shu-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func printNumbers(n int) []int {
	size := 1
	for i := 0; i < n; i++ {
		size *= 10
	}
	out := make([]int, size-1)
	for i := range out {
		out[i] = i + 1
	}
	return out
}
