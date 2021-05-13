package main

/*
  Given an input string (s) and a pattern (p), implement regular expression matching with support for '.' and '*' where:
    1. '.' Matches any single character.​​​​
    2. '*' Matches zero or more of the preceding element.

  The matching should cover the entire input string (not partial).

  Example 1:
    Input: s = "aa", p = "a"
    Output: false
    Explanation: "a" does not match the entire string "aa".

  Example 2:
    Input: s = "aa", p = "a*"
    Output: true
    Explanation: '*' means zero or more of the preceding element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".

  Example 3:
    Input: s = "ab", p = ".*"
    Output: true
    Explanation: ".*" means "zero or more (*) of any character (.)".

  Example 4:
    Input: s = "aab", p = "c*a*b"
    Output: true
    Explanation: c can be repeated 0 times, a can be repeated 1 time. Therefore, it matches "aab".

  Example 5:
    Input: s = "mississippi", p = "mis*is*p*."
    Output: false

  Constraints:
    1. 0 <= s.length <= 20
    2. 0 <= p.length <= 30
    3. s contains only lowercase English letters.
    4. p contains only lowercase English letters, '.', and '*'.
    5. It is guaranteed for each appearance of the character '*', there will be a previous valid character to match.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/regular-expression-matching
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Dynamic Programming
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	f := make([][]bool, m+1)
	for i := range f {
		f[i] = make([]bool, n+1)
	}

	f[0][0] = true
	for i := 1; i < n && f[0][i-1]; i += 2 {
		if p[i] == '*' {
			f[0][i+1] = true
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if s[i] == p[j] || '.' == p[j] {
				f[i+1][j+1] = f[i][j]
			} else if '*' == p[j] {
				f[i+1][j+1] = f[i+1][j-1] // repeat 0 times
				if s[i] == p[j-1] || '.' == p[j-1] {
					f[i+1][j+1] = f[i+1][j+1] || f[i+1][j] || f[i][j+1] // repeat 1 or more times
				}
			}
		}
	}
	return f[m][n]
}
