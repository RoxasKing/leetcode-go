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

// Stack
func calculate(s string) int {
	nus := []int{}  // number stack
	ops := []byte{} // operator stack
	for s != "" {
		for s != "" && s[0] == ' ' {
			s = s[1:]
		}
		if s == "" {
			break
		}
		if s[0] == '+' || s[0] == '-' || s[0] == '(' {
			ops = append(ops, s[0])
			s = s[1:]
		} else if s[0] == ')' {
			s = s[1:]
			if len(ops) > 0 && ops[len(ops)-1] != '(' {
				op := ops[len(ops)-1]
				ops = ops[:len(ops)-1]
				num := nus[len(nus)-1]
				nus = nus[:len(nus)-1]
				if op == '-' {
					num = -num
				}
				nus = append(nus, num)
			} else if len(ops) > 0 {
				ops = ops[:len(ops)-1]
			}
			if len(ops) > 0 && ops[len(ops)-1] != '(' {
				num2 := nus[len(nus)-1]
				nus = nus[:len(nus)-1]
				num1 := 0
				if len(nus) > 0 {
					num1 = nus[len(nus)-1]
					nus = nus[:len(nus)-1]
				}
				op := ops[len(ops)-1]
				ops = ops[:len(ops)-1]
				res := 0
				if op == '+' {
					res = num1 + num2
				} else {
					res = num1 - num2
				}
				nus = append(nus, res)
			} else if len(ops) > 0 {
				ops = ops[:len(ops)-1]
			}
		} else {
			num2 := 0
			for s != "" && '0' <= s[0] && s[0] <= '9' {
				num2 = num2*10 + int(s[0]-'0')
				s = s[1:]
			}
			if len(ops) > 0 && ops[len(ops)-1] != '(' {
				op := ops[len(ops)-1]
				ops = ops[:len(ops)-1]
				num1 := 0
				if len(nus) > 0 {
					num1 = nus[len(nus)-1]
					nus = nus[:len(nus)-1]
				}
				res := 0
				if op == '+' {
					res = num1 + num2
				} else {
					res = num1 - num2
				}
				nus = append(nus, res)
			} else {
				nus = append(nus, num2)
			}
		}
	}
	return nus[0]
}
