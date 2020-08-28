package main

/*
  给你一个字符串 s ，请你返回满足以下条件的最长子字符串的长度：每个元音字母，即 'a'，'e'，'i'，'o'，'u' ，在子字符串中都恰好出现了偶数次。

  提示：
    1 <= s.length <= 5 x 10^5
    s 只包含小写英文字母。

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/find-the-longest-substring-containing-vowels-in-even-counts
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func findTheLongestSubstring(s string) int {
	ans, status := 0, 0
	pos := make([]int, 1<<5)
	for i := 0; i < len(pos); i++ {
		pos[i] = -1
	}
	pos[0] = 0
	for i := 1; i <= len(s); i++ {
		switch s[i-1] {
		case 'a':
			status ^= 1 << 0
		case 'e':
			status ^= 1 << 1
		case 'i':
			status ^= 1 << 2
		case 'o':
			status ^= 1 << 3
		case 'u':
			status ^= 1 << 4
		}
		if pos[status] != -1 {
			ans = Max(ans, i-pos[status])
		} else {
			pos[status] = i
		}
	}
	return ans
}
