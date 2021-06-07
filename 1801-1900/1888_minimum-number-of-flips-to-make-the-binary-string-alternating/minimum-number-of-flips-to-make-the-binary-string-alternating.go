package main

/*
  You are given a binary string s. You are allowed to perform two types of operations on the string in any sequence:

    Type-1: Remove the character at the start of the string s and append it to the end of the string.
    Type-2: Pick any character in s and flip its value, i.e., if its value is '0' it becomes '1' and vice-versa.

  Return the minimum number of type-2 operations you need to perform such that s becomes alternating.

  The string is called alternating if no two adjacent characters are equal.

  For example, the strings "010" and "1010" are alternating, while the string "0100" is not.

  Example 1:
    Input: s = "111000"
    Output: 2
    Explanation: Use the first operation two times to make s = "100011".
      Then, use the second operation on the third and sixth elements to make s = "101010".

  Example 2:
    Input: s = "010"
    Output: 0
    Explanation: The string is already alternating.

  Example 3:
    Input: s = "1110"
    Output: 1
    Explanation: Use the second operation on the second element to make s = "1010".

  Constraints:
    1. 1 <= s.length <= 10^5
    2. s[i] is either '0' or '1'.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/minimum-number-of-flips-to-make-the-binary-string-alternating
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Prefix Sum
// Greedy Algorithm

func minFlips(s string) int {
	n := len(s)
	pSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		pSum[i+1] = pSum[i] + (int(s[i]-'0') ^ (i & 1))
	}

	out := Min(pSum[n], n-pSum[n])
	for i := 1; i < n; i++ {
		cnt := 0
		if i&1 == 0 {
			cnt += pSum[n] - pSum[i]
		} else {
			cnt += (n - i) - (pSum[n] - pSum[i])
		}
		if (n-i)&1 == 0 {
			cnt += pSum[i]
		} else {
			cnt += i - pSum[i]
		}
		out = Min(out, Min(cnt, n-cnt))
	}
	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
