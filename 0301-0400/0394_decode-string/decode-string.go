package main

import (
	"strconv"
	"strings"
)

/*
  Given an encoded string, return its decoded string.

  The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets is being repeated exactly k times. Note that k is guaranteed to be a positive integer.

  You may assume that the input string is always valid; No extra white spaces, square brackets are well-formed, etc.

  Furthermore, you may assume that the original data does not contain any digits and that digits are only for those repeat numbers, k. For example, there won't be input like 3a or 2[4].

  Example 1:
    Input: s = "3[a]2[bc]"
    Output: "aaabcbc"

  Example 2:
    Input: s = "3[a2[c]]"
    Output: "accaccacc"

  Example 3:
    Input: s = "2[abc]3[cd]ef"
    Output: "abcabccdcdcdef"

  Example 4:
    Input: s = "abc3[cd]xyz"
    Output: "abccdcdcdxyz"

  Constraints:
    1. 1 <= s.length <= 30
    2. s consists of lowercase English letters, digits, and square brackets '[]'.
    3. s is guaranteed to be a valid input.
    4. All the integers in s are in the range [1, 300].

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/decode-string
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Stack
func decodeString(s string) string {
	stack := []string{}
	for r := 0; r < len(s); {
		if '0' <= s[r] && s[r] <= '9' {
			l := r
			for '0' <= s[r] && s[r] <= '9' {
				r++
			}
			stack = append(stack, s[l:r])
		} else if s[r] == ']' {
			tmp := []string{}
			for stack[len(stack)-1] != "[" {
				tmp = append(tmp, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			for i := 0; i < len(tmp)/2; i++ {
				tmp[i], tmp[len(tmp)-i-1] = tmp[len(tmp)-i-1], tmp[i]
			}
			stack = stack[:len(stack)-1] // delete "["
			times, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-1]
			stack = append(stack, strings.Repeat(strings.Join(tmp, ""), times))
			r++
		} else {
			stack = append(stack, s[r:r+1])
			r++
		}
	}
	return strings.Join(stack, "")
}
