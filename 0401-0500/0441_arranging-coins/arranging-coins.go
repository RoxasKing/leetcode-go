package main

/*
  You have n coins and you want to build a staircase with these coins. The staircase consists of k rows where the ith row has exactly i coins. The last row of the staircase may be incomplete.

  Given the integer n, return the number of complete rows of the staircase you will build.

  Example 1:
    Input: n = 5
    Output: 2
    Explanation: Because the 3rd row is incomplete, we return 2.

  Example 2:
    Input: n = 8
    Output: 3
    Explanation: Because the 4th row is incomplete, we return 3.

  Constraints:
    1 <= n <= 2^31 - 1

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/arranging-coins
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Binary Search
func arrangeCoins(n int) int {
	l, r := 0, n+1
	for l < r {
		m := (l + r) >> 1
		if m*(1+m)>>1 <= n {
			l = m + 1
		} else {
			r = m
		}
	}
	return l - 1
}
