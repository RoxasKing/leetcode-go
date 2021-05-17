package main

/*
  There are n uniquely-sized sticks whose lengths are integers from 1 to n. You want to arrange the sticks such that exactly k sticks are visible from the left. A stick is visible from the left if there are no longer sticks to the left of it.

    For example, if the sticks are arranged [1,3,2,5,4], then the sticks with lengths 1, 3, and 5 are visible from the left.

  Given n and k, return the number of such arrangements. Since the answer may be large, return it modulo 109 + 7.

  Example 1:
    Input: n = 3, k = 2
    Output: 3
    Explanation: [1,3,2], [2,3,1], and [2,1,3] are the only arrangements such that exactly 2 sticks are visible.
      The visible sticks are underlined.

  Example 2:
    Input: n = 5, k = 5
    Output: 1
    Explanation: [1,2,3,4,5] is the only arrangement such that all 5 sticks are visible.
      The visible sticks are underlined.

  Example 3:
    Input: n = 20, k = 11
    Output: 647427950
    Explanation: There are 647427950 (mod 109 + 7) ways to rearrange the sticks such that exactly 11 sticks are visible.

  Constraints:
    1. 1 <= n <= 1000
    2. 1 <= k <= n

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/number-of-ways-to-rearrange-sticks-with-k-sticks-visible
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Dynamic Programming

func rearrangeSticks(n int, k int) int {
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, k+1)
	}
	f[0][0] = 1

	for i := 1; i <= n; i++ {
		for j := 1; j <= k && j <= i; j++ {
			f[i][j] = f[i-1][j]*(i-1) + f[i-1][j-1]
			f[i][j] %= 1e9 + 7
		}
	}

	return f[n][k]
}
