package main

/*
  Given an integer n, return the number of strings of length n that consist only of vowels (a, e, i, o, u) and are lexicographically sorted.

  A string s is lexicographically sorted if for all valid i, s[i] is the same as or comes before s[i+1] in the alphabet.

  Example 1:
    Input: n = 1
    Output: 5
    Explanation: The 5 sorted strings that consist of vowels only are ["a","e","i","o","u"].

  Example 2:
    Input: n = 2
    Output: 15
    Explanation: The 15 sorted strings that consist of vowels only are
    ["aa","ae","ai","ao","au","ee","ei","eo","eu","ii","io","iu","oo","ou","uu"].
    Note that "ea" is not a valid string since 'e' comes after 'a' in the alphabet.

  Example 3:
    Input: n = 33
    Output: 66045

  Constraints:
    1 <= n <= 50

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/count-sorted-vowel-strings
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Dynamic Programming
func countVowelStrings(n int) int {
	dp := [2][5]int{}
	for i := 0; i < 5; i++ {
		dp[0][i] = i + 1
	}
	for i := 1; i < n; i++ {
		dp[1][0] = 1
		for j := 1; j < 5; j++ {
			dp[1][j] = dp[1][j-1] + dp[0][j]
		}
		dp[0], dp[1] = dp[1], dp[0]
	}
	return dp[0][4]
}
