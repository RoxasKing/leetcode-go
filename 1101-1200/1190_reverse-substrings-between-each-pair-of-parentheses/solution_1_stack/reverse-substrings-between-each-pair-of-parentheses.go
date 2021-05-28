package main

/*
  You are given a string s that consists of lower case English letters and brackets.

  Reverse the strings in each pair of matching parentheses, starting from the innermost one.

  Your result should not contain any brackets.

  Example 1:
    Input: s = "(abcd)"
    Output: "dcba"

  Example 2:
    Input: s = "(u(love)i)"
    Output: "iloveu"
    Explanation: The substring "love" is reversed first, then the whole string is reversed.

  Example 3:
    Input: s = "(ed(et(oc))el)"
    Output: "leetcode"
    Explanation: First, we reverse the substring "oc", then "etco", and finally, the whole string.

  Example 4:
    Input: s = "a(bcdefghijkl(mno)p)q"
    Output: "apmnolkjihgfedcbq"

  Constraints:
    1. 0 <= s.length <= 2000
    2. s only contains lower case English characters and parentheses.
    3. It's guaranteed that all parentheses are balanced.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/reverse-substrings-between-each-pair-of-parentheses
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Stack

func reverseParentheses(s string) string {
	stk := RuneStack{}
	for _, ch := range s {
		if ch == ')' {
			tmp := []rune{}
			for stk.Top() != '(' {
				tmp = append(tmp, stk.Pop())
			}
			stk.Pop()
			stk.Push(tmp...)
		} else {
			stk.Push(ch)
		}
	}
	return string(stk)
}

type RuneStack []rune

func (s RuneStack) Len() int        { return len(s) }
func (s RuneStack) Top() rune       { return s[s.Len()-1] }
func (s *RuneStack) Push(x ...rune) { *s = append(*s, x...) }
func (s *RuneStack) Pop() rune {
	top := s.Len() - 1
	out := (*s)[top]
	*s = (*s)[:top]
	return out
}
