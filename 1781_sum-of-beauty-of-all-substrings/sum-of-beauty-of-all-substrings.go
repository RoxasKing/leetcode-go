package main

/*
  The beauty of a string is the difference in frequencies between the most frequent and least frequent characters.
    For example, the beauty of "abaacc" is 3 - 1 = 2.
  Given a string s, return the sum of beauty of all of its substrings.

  Example 1:
    Input: s = "aabcb"
    Output: 5
    Explanation: The substrings with non-zero beauty are ["aab","aabc","aabcb","abcb","bcb"], each with beauty equal to 1.

  Example 2:
    Input: s = "aabcbaa"
    Output: 17

  Constraints:
    1. 1 <= s.length <= 500
    2. s consists of only lowercase English letters.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/sum-of-beauty-of-all-substrings
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Sliding Window
func beautySum(s string) int {
	out := 0
	for k := 3; k <= len(s); k++ {
		cnt := [26]int{}
		for i := range s {
			cnt[s[i]-'a']++
			if i > k-1 {
				cnt[s[i-k]-'a']--
			}
			if i >= k-1 {
				min, max := 1<<31-1, -1<<31
				for j := 0; j < 26; j++ {
					if cnt[j] > 0 {
						min = Min(min, cnt[j])
						max = Max(max, cnt[j])
					}
				}
				out += max - min
			}
		}
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
