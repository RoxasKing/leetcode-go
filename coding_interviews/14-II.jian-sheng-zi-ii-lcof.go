package codinginterviews

/*
  给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为 k[0],k[1]...k[m] 。请问 k[0]*k[1]*...*k[m] 可能的最大乘积是多少？例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。
  答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/jian-sheng-zi-ii-lcof
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func cuttingRopeII(n int) int {
	pow := func(a, b int) int {
		res := 1
		for i := 0; i < b; i++ {
			res = (res * a) % (1e9 + 7)
		}
		return res
	}
	if n < 2 {
		return 0
	} else if n == 2 {
		return 1
	} else if n == 3 {
		return 2
	}
	timesOf3 := n / 3
	if n-timesOf3*3 == 1 {
		timesOf3--
	}
	timesOf2 := (n - timesOf3*3) >> 1
	return (pow(3, timesOf3) * pow(2, timesOf2)) % (1e9 + 7)
}
