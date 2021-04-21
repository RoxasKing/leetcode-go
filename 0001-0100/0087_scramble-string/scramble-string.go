package main

/*
  We can scramble a string s to get a string t using the following algorithm:
    1. If the length of the string is 1, stop.
    2. If the length of the string is > 1, do the following:
       1. Split the string into two non-empty substrings at a random index, i.e., if the string is s, divide it to x and y where s = x + y.
       2. Randomly decide to swap the two substrings or to keep them in the same order. i.e., after this step, s may become s = x + y or s = y + x.
	   3. Apply step 1 recursively on each of the two substrings x and y.
  Given two strings s1 and s2 of the same length, return true if s2 is a scrambled string of s1, otherwise, return false.

  Example 1:
    Input: s1 = "great", s2 = "rgeat"
    Output: true
    Explanation: One possible scenario applied on s1 is:
      "great" --> "gr/eat" // divide at random index.
      "gr/eat" --> "gr/eat" // random decision is not to swap the two substrings and keep them in order.
      "gr/eat" --> "g/r / e/at" // apply the same algorithm recursively on both substrings. divide at ranom index each of them.
      "g/r / e/at" --> "r/g / e/at" // random decision was to swap the first substring and to keep the second substring in the same order.
      "r/g / e/at" --> "r/g / e/ a/t" // again apply the algorithm recursively, divide "at" to "a/t".
      "r/g / e/ a/t" --> "r/g / e/ a/t" // random decision is to keep both substrings in the same order.
      The algorithm stops now and the result string is "rgeat" which is s2.
      As there is one possible scenario that led s1 to be scrambled to s2, we return true.

  Example 2:
    Input: s1 = "abcde", s2 = "caebd"
    Output: false

  Example 3:
    Input: s1 = "a", s2 = "a"
    Output: true

  Constraints:
    1. s1.length == s2.length
    2. 1 <= s1.length <= 30
    3. s1 and s2 consist of lower-case English letters.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/scramble-string
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Dynamic Programming + DFS
func isScramble(s1 string, s2 string) bool {
	valid := make(map[string]bool)
	return dp(valid, s1, s2)
}

func dp(valid map[string]bool, s1, s2 string) bool {
	if s1 > s2 {
		s1, s2 = s2, s1
	}
	key := s1 + s2
	if val, ok := valid[key]; ok {
		return val
	}
	if s1 == s2 {
		valid[key] = true
		return true
	}
	cnt := make([]int, 128)
	for i := range s1 {
		cnt[s1[i]]++
	}
	for i := range s2 {
		if cnt[s2[i]]--; cnt[s2[i]] < 0 {
			valid[key] = false
			return false
		}
	}
	n := len(s1)
	for i := 1; i < n; i++ {
		if dp(valid, s1[:i], s2[:i]) && dp(valid, s1[i:], s2[i:]) ||
			dp(valid, s1[:i], s2[n-i:]) && dp(valid, s1[i:], s2[:n-i]) {
			valid[key] = true
			return true
		}
	}
	valid[key] = false
	return false
}
