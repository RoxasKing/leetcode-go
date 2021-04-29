package main

/*
  Given a string s representing an expression, implement a basic calculator to evaluate it.

  Example 1:
    Input: s = "1 + 1"
    Output: 2

  Example 2:
    Input: s = " 2-1 + 2 "
    Output: 3

  Example 3:
    Input: s = "(1+(4+5+2)-3)+(6+8)"
    Output: 23

  Constraints:
    1. 1 <= s.length <= 3 * 105
    2. s consists of digits, '+', '-', '(', ')', and ' '.
    3. s represents a valid expression.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/basic-calculator
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Stack
func calculate(s string) int {
	out := 0
	ops := []int{1}
	sig := 1
	n := len(s)
	for i := 0; i < n; {
		switch s[i] {
		case ' ':
			i++
		case '+':
			sig = ops[len(ops)-1]
			i++
		case '-':
			sig = -ops[len(ops)-1]
			i++
		case '(':
			ops = append(ops, sig)
			i++
		case ')':
			ops = ops[:len(ops)-1]
			i++
		default:
			num := 0
			for ; i < n && '0' <= s[i] && s[i] <= '9'; i++ {
				num = num*10 + int(s[i]-'0')
			}
			out += sig * num
		}
	}
	return out
}
