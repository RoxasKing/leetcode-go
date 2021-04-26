package main

/*
  Given an integer n (in base 10) and a base k, return the sum of the digits of n after converting n from base 10 to base k.

  After converting, each digit should be interpreted as a base 10 number, and the sum should be returned in base 10.

  Example 1:
    Input: n = 34, k = 6
    Output: 9
    Explanation: 34 (base 10) expressed in base 6 is 54. 5 + 4 = 9.

  Example 2:
    Input: n = 10, k = 10
    Output: 1
    Explanation: n is already in base 10. 1 + 0 = 1.

  Constraints:
    1. 1 <= n <= 100
    2. 2 <= k <= 10

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/sum-of-digits-in-base-k
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func sumBase(n int, k int) int {
	out := 0
	for n > 0 {
		out += n % k
		n /= k
	}
	return out
}
