package leetcode

/*
  我们正在玩一个猜数字游戏。 游戏规则如下：
  我从 1 到 n 选择一个数字。 你需要猜我选择了哪个数字。
  每次你猜错了，我会告诉你这个数字是大了还是小了。
  你调用一个预先定义好的接口 guess(int num)，它会返回 3 个可能的结果（-1，1 或 0）：

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/guess-number-higher-or-lower
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	l, r := 1, n
	for l <= r {
		m := l + (r-l)>>1
		res := guess(m)
		if res == -1 {
			r = m - 1
		} else if res == 1 {
			l = m + 1
		} else {
			return m
		}
	}
	return -1
}

func guess(num int) int {
	if num < 6 {
		return -1
	} else if num > 6 {
		return 1
	}
	return 0
}
