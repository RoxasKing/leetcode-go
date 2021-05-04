package main

/*
  Given a string s which represents an expression, evaluate this expression and return its value.

  The integer division should truncate toward zero.

  Example 1:
    Input: s = "3+2*2"
    Output: 7

  Example 2:
    Input: s = " 3/2 "
    Output: 1

  Example 3:
    Input: s = " 3+5 / 2 "
    Output: 5

  Constraints:
    1. 1 <= s.length <= 3 * 10^5
    2. s consists of integers and operators ('+', '-', '*', '/') separated by some number of spaces.
    3. s represents a valid expression.
    4. All the integers in the expression are non-negative integers in the range [0, 231 - 1].
    5. The answer is guaranteed to fit in a 32-bit integer.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/basic-calculator-ii
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Stack
func calculate(s string) int {
	stk := IntStack{}
	num := 0
	op := '+'

	for _, ch := range s + "+" {
		if ch == ' ' {
			continue
		}
		if '0' <= ch && ch <= '9' {
			num = num*10 + int(ch-'0')
			continue
		}
		switch op {
		case '+':
			stk.Push(num)
		case '-':
			stk.Push(-num)
		case '*':
			stk.Push(stk.Pop() * num)
		case '/':
			stk.Push(stk.Pop() / num)
		}
		num = 0
		op = ch
	}

	out := 0
	for _, num := range stk {
		out += num
	}
	return out
}

type IntStack []int

func (s IntStack) Len() int    { return len(s) }
func (s IntStack) Top() int    { return s[s.Len()-1] }
func (s *IntStack) Push(x int) { *s = append(*s, x) }
func (s *IntStack) Pop() int {
	top := s.Len() - 1
	out := (*s)[top]
	*s = (*s)[:top]
	return out
}
