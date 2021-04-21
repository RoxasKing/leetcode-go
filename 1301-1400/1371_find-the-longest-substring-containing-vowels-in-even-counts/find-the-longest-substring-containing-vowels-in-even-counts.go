package main

/*
  Given the string s, return the size of the longest substring containing each vowel an even number of times. That is, 'a', 'e', 'i', 'o', and 'u' must appear an even number of times.

  Example 1:
    Input: s = "eleetminicoworoep"
    Output: 13
    Explanation: The longest substring is "leetminicowor" which contains two each of the vowels: e, i and o and zero of the vowels: a and u.

  Example 2:
    Input: s = "leetcodeisgreat"
    Output: 5
    Explanation: The longest substring is "leetc" which contains two e's.

  Example 3:
    Input: s = "bcbcbc"
    Output: 6
    Explanation: In this case, the given string "bcbcbc" is the longest because all vowels: a, e, i, o and u appear zero times.

  Constraints:
    1. 1 <= s.length <= 5 x 10^5
    2. s contains only lowercase English letters.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/find-the-longest-substring-containing-vowels-in-even-counts
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Prefix Sum + Bit Manipulation
func findTheLongestSubstring(s string) int {
	pos := make([]int, 1<<5)
	pos[0] = -1
	for i := 1; i < 1<<5; i++ {
		pos[i] = -2
	}
	out, xor := 0, 0
	for i := range s {
		switch s[i] {
		case 'a':
			xor ^= 1 << 0
		case 'e':
			xor ^= 1 << 1
		case 'i':
			xor ^= 1 << 2
		case 'o':
			xor ^= 1 << 3
		case 'u':
			xor ^= 1 << 4
		}
		if j := pos[xor]; j == -2 {
			pos[xor] = i
		} else if i-j > out {
			out = i - j
		}
	}
	return out
}
