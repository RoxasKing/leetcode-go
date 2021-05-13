package main

/*
  You have a pointer at index 0 in an array of size arrLen. At each step, you can move 1 position to the left, 1 position to the right in the array or stay in the same place  (The pointer should not be placed outside the array at any time).

  Given two integers steps and arrLen, return the number of ways such that your pointer still at index 0 after exactly steps steps.

  Since the answer may be too large, return it modulo 10^9 + 7.

  Example 1:
    Input: steps = 3, arrLen = 2
    Output: 4
    Explanation: There are 4 differents ways to stay at index 0 after 3 steps.
      Right, Left, Stay
      Stay, Right, Left
      Right, Stay, Left
      Stay, Stay, Stay

  Example 2:
    Input: steps = 2, arrLen = 4
    Output: 2
    Explanation: There are 2 differents ways to stay at index 0 after 2 steps
      Right, Left
      Stay, Stay

  Example 3:
    Input: steps = 4, arrLen = 2
    Output: 8

  Constraints:
    1. 1 <= steps <= 500
    2. 1 <= arrLen <= 10^6

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/number-of-ways-to-stay-in-the-same-place-after-some-steps
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Dynamic Programming
func numWays(steps int, arrLen int) int {
	arrLen = Min(arrLen, steps+1)
	f := make([][]int, steps+1)
	for i := range f {
		f[i] = make([]int, arrLen)
	}
	f[0][0] = 1
	for i := 1; i <= steps; i++ {
		for j := 0; j < arrLen; j++ {
			f[i][j] = f[i-1][j]
			if j-1 >= 0 {
				f[i][j] += f[i-1][j-1]
				f[i][j] %= 1e9 + 7
			}
			if j+1 < arrLen {
				f[i][j] += f[i-1][j+1]
				f[i][j] %= 1e9 + 7
			}
		}
	}
	return f[steps][0]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
