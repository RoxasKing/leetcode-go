package main

/*
  A valid number can be split up into these components (in order):

    1. A decimal number or an integer.
    2. (Optional) An 'e' or 'E', followed by an integer.

  A decimal number can be split up into these components (in order):

    1. (Optional) A sign character (either '+' or '-').
    2. One of the following formats:
       1. One or more digits, followed by a dot '.'.
       2. One or more digits, followed by a dot '.', followed by one or more digits.
       3. A dot '.', followed by one or more digits.

  An integer can be split up into these components (in order):

  (Optional) A sign character (either '+' or '-').
  One or more digits.
  For example, all the following are valid numbers: ["2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789"], while the following are not valid numbers: ["abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53"].

  Given a string s, return true if s is a valid number.

  Example 1:
    Input: s = "0"
    Output: true

  Example 2:
    Input: s = "e"
    Output: false

  Example 3:
    Input: s = "."
    Output: false

  Example 4:
    Input: s = ".1"
    Output: true

  Constraints:
    1. 1 <= s.length <= 20
    2. s consists of only English letters (both uppercase and lowercase), digits (0-9), plus '+', minus '-', or dot '.'.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/valid-number
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

func isNumber(s string) bool {
	trimSpace(&s)
	havePrefixNum := walkInteger(&s)
	if len(s) > 0 && s[0] == '.' {
		s = s[1:]
		if ok := walkPureInteger(&s); !ok && !havePrefixNum {
			return false
		}
		havePrefixNum = true
	}
	if len(s) > 0 && (s[0] == 'E' || s[0] == 'e') {
		if !havePrefixNum {
			return false
		}
		if s = s[1:]; !walkInteger(&s) {
			return false
		}
	}
	return len(s) == 0
}

func trimSpace(s *string) {
	for len(*s) > 0 && (*s)[0] == ' ' {
		*s = (*s)[1:]
	}
	for len(*s) > 0 && (*s)[len(*s)-1] == ' ' {
		*s = (*s)[:len(*s)-1]
	}
}

func walkInteger(s *string) bool {
	if len(*s) > 0 && ((*s)[0] == '+' || (*s)[0] == '-') {
		*s = (*s)[1:]
	}
	return walkPureInteger(s)
}

func walkPureInteger(s *string) bool {
	cnt := 0
	for len(*s) > 0 && '0' <= (*s)[0] && (*s)[0] <= '9' {
		*s = (*s)[1:]
		cnt++
	}
	return cnt > 0
}
