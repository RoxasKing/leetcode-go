package main

import (
	"fmt"
)

func getPermutation(n int, k int) string {
	if n < 1 {
		return ""
	}
	// 存放结果字符集
	out := make([]byte, n)
	// 存放待取字符集
	tmp := make([]byte, n)
	for i := 0; i < n; i++ {
		// 初始化待取结果集
		tmp[i] = byte(i + 1 + '0')
	}
	// 数组最大下标为 len-1, 故 --
	k--

	// 获取最大基数,即n-1时的排列数量
	base := 1
	for i := 2; i < n; i++ {
		base *= i
	}

	for i := 0; i < n-1; i++ {
		// 通过目标排列除以基数，计算出需要取出的字符的下标
		index := k / base
		out[i] = tmp[index]
		// 将已取出字符后面字符全部往前赶
		tmp = append(tmp[:index], tmp[index+1:]...)
		// 下一个目标排列
		k %= base
		// 下一个基数
		base /= (n - i - 1)
	}
	// 最后一个字符
	out[n-1] = tmp[0]
	return string(out)
}

func main() {
	fmt.Println(getPermutation(3, 6))
}
