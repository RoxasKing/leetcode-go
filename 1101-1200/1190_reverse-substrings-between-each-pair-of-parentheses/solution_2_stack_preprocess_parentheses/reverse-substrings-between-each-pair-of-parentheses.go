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

// Important!

// Stack
/* preprocess parentheses */

func reverseParentheses(s string) string {
	n := len(s)
	pair := make([]int, n)
	stk := IntStack{}
	for i, b := range s {
		if b == '(' {
			stk.Push(i)
		} else if b == ')' {
			j := stk.Pop()
			pair[i], pair[j] = j, i
		}
	}

	out := []byte{}
	for i, step := 0, 1; i < n; i += step {
		if s[i] == '(' || s[i] == ')' {
			i = pair[i]
			step = -step
		} else {
			out = append(out, s[i])
		}
	}
	return string(out)
}

type IntStack []int

func (s IntStack) Len() int       { return len(s) }
func (s IntStack) Top() int       { return s[s.Len()-1] }
func (s *IntStack) Push(x ...int) { *s = append(*s, x...) }
func (s *IntStack) Pop() int {
	top := s.Len() - 1
	out := (*s)[top]
	*s = (*s)[:top]
	return out
}
