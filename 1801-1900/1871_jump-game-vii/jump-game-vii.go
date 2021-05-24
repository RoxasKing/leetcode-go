package main

/*
  You are given a 0-indexed binary string s and two integers minJump and maxJump. In the beginning, you are standing at index 0, which is equal to '0'. You can move from index i to index j if the following conditions are fulfilled:

    1. i + minJump <= j <= min(i + maxJump, s.length - 1), and
    2. s[j] == '0'.

  Return true if you can reach index s.length - 1 in s, or false otherwise.

  Example 1:
    Input: s = "011010", minJump = 2, maxJump = 3
    Output: true
    Explanation:
      In the first step, move from index 0 to index 3.
      In the second step, move from index 3 to index 5.

  Example 2:
    Input: s = "01101110", minJump = 2, maxJump = 3
    Output: false

  Constraints:
    1. 2 <= s.length <= 10^5
    2. s[i] is either '0' or '1'.
    3. s[0] == '0'
    4. 1 <= minJump <= maxJump < s.length

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/jump-game-vii
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Greedy Algoritym

func canReach(s string, minJump int, maxJump int) bool {
	n := len(s)
	if s[n-1] != '0' {
		return false
	}

	l, r := minJump, Min(maxJump, n-1)
	for l <= r && r < n-1 {
		nextL, nextR := 1<<31-1, -1<<31
		for i := l; i <= Min(r, n-1); i++ {
			if s[i] == '0' {
				nextL, nextR = Min(nextL, i+minJump), i+maxJump
			}
		}
		l, r = Max(r+1, nextL), nextR
	}
	return l <= n-1 && n-1 <= r
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
