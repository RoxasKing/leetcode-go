package main

/*
  You are given a string s (0-indexed)​​​​​​. You are asked to perform the following operation on s​​​​​​ until you get a sorted string:
    1. Find the largest index i such that 1 <= i < s.length and s[i] < s[i - 1].
    2. Find the largest index j such that i <= j < s.length and s[k] < s[i - 1] for all the possible values of k in the range [i, j] inclusive.
    3. Swap the two characters at indices i - 1​​​​ and j​​​​​.
    4. Reverse the suffix starting at index i​​​​​​.
  Return the number of operations needed to make the string sorted. Since the answer can be too large, return it modulo 109 + 7.

  Example 1:
    Input: s = "cba"
    Output: 5
    Explanation: The simulation goes as follows:
      Operation 1: i=2, j=2. Swap s[1] and s[2] to get s="cab", then reverse the suffix starting at 2. Now, s="cab".
      Operation 2: i=1, j=2. Swap s[0] and s[2] to get s="bac", then reverse the suffix starting at 1. Now, s="bca".
      Operation 3: i=2, j=2. Swap s[1] and s[2] to get s="bac", then reverse the suffix starting at 2. Now, s="bac".
      Operation 4: i=1, j=1. Swap s[0] and s[1] to get s="abc", then reverse the suffix starting at 1. Now, s="acb".
      Operation 5: i=2, j=2. Swap s[1] and s[2] to get s="abc", then reverse the suffix starting at 2. Now, s="abc".

  Example 2:
    Input: s = "aabaa"
    Output: 2
    Explanation: The simulation goes as follows:
      Operation 1: i=3, j=4. Swap s[2] and s[4] to get s="aaaab", then reverse the substring starting at 3. Now, s="aaaba".
      Operation 2: i=4, j=4. Swap s[3] and s[4] to get s="aaaab", then reverse the substring starting at 4. Now, s="aaaab".

  Example 3:
    Input: s = "cdbea"
    Output: 63

  Example 4:
    Input: s = "leetcodeleetcodeleetcode"
    Output: 982157772

  Constraints:
    1. 1 <= s.length <= 3000
    k. s​​​​​​ consists only of lowercase English letters.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/minimum-number-of-operations-to-make-string-sorted
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Math
func makeStringSorted(s string) int {
	n := len(s)
	fac := make([]int, n+1)
	fac[0] = 1
	facinv := make([]int, n+1)
	facinv[0] = 1
	for i := 1; i < n; i++ {
		fac[i] = fac[i-1] * i % mod
		facinv[i] = quickMul(fac[i], mod-2) // 费马小定理
	}

	freq := [26]int{}
	for i := range s {
		freq[s[i]-'a']++
	}

	out := 0
	for i := 0; i < n-1; i++ {
		rank := 0
		idx := int(s[i] - 'a')
		for j := 0; j < idx; j++ {
			rank += freq[j]
		}
		cnt := rank * fac[n-1-i] % mod
		for j := 0; j < 26; j++ {
			cnt *= facinv[freq[j]]
			cnt %= mod
		}
		out += cnt
		out %= mod
		freq[idx]--
	}
	return out
}

func quickMul(x, n int) int {
	out := 1
	for n > 0 {
		if n&1 == 1 {
			out *= x
			out %= mod
		}
		x *= x
		x %= mod
		n >>= 1
	}
	return out
}

var mod = int(1e9 + 7)
