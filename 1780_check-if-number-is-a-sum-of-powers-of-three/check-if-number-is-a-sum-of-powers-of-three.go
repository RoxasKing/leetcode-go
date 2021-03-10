package main

/*
  Given an integer n, return true if it is possible to represent n as the sum of distinct powers of three. Otherwise, return false.

  An integer y is a power of three if there exists an integer x such that y == 3x.

  Example 1:
    Input: n = 12
    Output: true
    Explanation: 12 = 31 + 32

  Example 2:
    Input: n = 91
    Output: true
    Explanation: 91 = 30 + 32 + 34

  Example 3:
    Input: n = 21
    Output: false

  Constraints:
    1 <= n <= 10^7

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/check-if-number-is-a-sum-of-powers-of-three
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Math
func checkPowersOfThree(n int) bool {
	for n > 0 {
		remain := n % 3
		if remain != 0 && remain != 1 {
			return false
		}
		n /= 3
	}
	return true
}
