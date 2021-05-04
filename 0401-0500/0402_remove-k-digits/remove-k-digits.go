package main

/*
  Given string num representing a non-negative integer num, and an integer k, return the smallest possible integer after removing k digits from num.

  Example 1:
    Input: num = "1432219", k = 3
    Output: "1219"
    Explanation: Remove the three digits 4, 3, and 2 to form the new number 1219 which is the smallest.

  Example 2:
    Input: num = "10200", k = 1
    Output: "200"
    Explanation: Remove the leading 1 and the number is 200. Note that the output must not contain leading zeroes.

  Example 3:
    Input: num = "10", k = 2
    Output: "0"
    Explanation: Remove all the digits from the number and it is left with nothing which is 0.

  Constraints:
    1. 1 <= k <= num.length <= 10^5
    2. num consists of only digits.
    3. num does not have any leading zeros except for the zero itself.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/remove-k-digits
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Monotone Stack
func removeKdigits(num string, k int) string {
	stk := CharStack{}

	for i := range num {
		for ; stk.Len() > 0 && stk.Top() > num[i] && k > 0; k-- {
			stk.Pop()
		}
		stk.Push(num[i])
	}

	for ; k > 0; k-- {
		stk.Pop()
	}

	for stk.Len() > 0 && stk[0] == '0' {
		stk = stk[1:]
	}

	if len(stk) == 0 {
		return "0"
	}
	return string(stk)
}

type CharStack []byte

func (s CharStack) Len() int     { return len(s) }
func (s CharStack) Top() byte    { return s[s.Len()-1] }
func (s *CharStack) Push(x byte) { *s = append(*s, x) }
func (s *CharStack) Pop() byte {
	top := s.Len() - 1
	out := (*s)[top]
	*s = (*s)[:top]
	return out
}
