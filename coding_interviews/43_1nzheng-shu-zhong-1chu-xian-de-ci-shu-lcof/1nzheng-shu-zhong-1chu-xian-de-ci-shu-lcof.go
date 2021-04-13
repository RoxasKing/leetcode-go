package main

/*
  输入一个整数 n ，求1～n这n个整数的十进制表示中1出现的次数。

  例如，输入12，1～12这些整数中包含1 的数字有1、10、11和12，1一共出现了5次。

  示例 1：
    输入：n = 12
    输出：5

  示例 2：
    输入：n = 13
    输出：6

  限制：
    1 <= n < 2^31

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/1nzheng-shu-zhong-1chu-xian-de-ci-shu-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Math
func countDigitOne(n int) int {
	out := 0
	for i := 1; i <= n; i *= 10 {
		div := i * 10
		out += (n / div) * i // e.g. if i == 10, (1234/100)*10 = 12*10, left side 0~11, right side 0~9
		remain := n % div
		if remain >= i<<1-1 { // e.g. if i == 10, 1234%100 = 34 >= 19, so 1210 ~ 1219 have i number
			out += i
		} else if remain >= i { // e.g. if i == 10, 1215%100 = 15 >= 10, so 1210 ~ 1215 have 15%i+1 = 6 number
			out += remain%i + 1
		}
	}
	return out
}
