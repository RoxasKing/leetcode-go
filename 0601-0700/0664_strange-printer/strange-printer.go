package main

/*
  There is a strange printer with the following two special properties:

    1. The printer can only print a sequence of the same character each time.
    2. At each turn, the printer can print new characters starting from and ending at any place and will cover the original existing characters.

  Given a string s, return the minimum number of turns the printer needed to print it.

  Example 1:
    Input: s = "aaabbb"
    Output: 2
    Explanation: Print "aaa" first and then print "bbb".

  Example 2:
    Input: s = "aba"
    Output: 2
    Explanation: Print "aaa" first and then print "b" from the second place of the string, which will cover the existing character 'a'.

  Constraints:
    1. 1 <= s.length <= 100
    2. s consists of lowercase English letters.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/strange-printer
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Dynamic Programming

func strangePrinter(s string) int {
	n := len(s)
	f := make([][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([]int, n)
		f[i][i] = 1
	}
	for k := 2; k <= n; k++ {
		for i := 0; i+(k-1) < n; i++ {
			f[i][i+k-1] = k
			for j := i + 1; j < i+k; j++ {
				tmp := f[i][j-1] + f[j][i+k-1]
				if s[i] == s[i+k-1] {
					tmp--
				}
				f[i][i+k-1] = Min(f[i][i+k-1], tmp)
			}
		}
	}
	return f[0][n-1]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
