package codinginterviews

/*
  输入一个整数 n ，求1～n这n个整数的十进制表示中1出现的次数。
  例如，输入12，1～12这些整数中包含1 的数字有1、10、11和12，1一共出现了5次。

  限制：
    1 <= n < 2^31

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/1nzheng-shu-zhong-1chu-xian-de-ci-shu-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func countDigitOne(n int) int {
	var out int
	for i := 1; i <= n; i *= 10 {
		divider := i * 10
		out += (n/divider)*i + Min(Max(n%divider-i+1, 0), i)
	}
	return out
}
