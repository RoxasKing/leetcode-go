package main

/*
  You are given a valid boolean expression as a string expression consisting of the characters '1','0','&' (bitwise AND operator),'|' (bitwise OR operator),'(', and ')'.

    For example, "()1|1" and "(1)&()" are not valid while "1", "(((1))|(0))", and "1|(0&(1))" are valid expressions.

  Return the minimum cost to change the final value of the expression.

    For example, if expression = "1|1|(0&0)&1", its value is 1|1|(0&0)&1 = 1|1|0&1 = 1|0&1 = 1&1 = 1. We want to apply operations so that the new expression evaluates to 0.

  The cost of changing the final value of an expression is the number of operations performed on the expression. The types of operations are described as follows:

    Turn a '1' into a '0'.
    Turn a '0' into a '1'.
    Turn a '&' into a '|'.
    Turn a '|' into a '&'.

  Note: '&' does not take precedence over '|' in the order of calculation. Evaluate parentheses first, then in left-to-right order.

  Example 1:
    Input: expression = "1&(0|1)"
    Output: 1
    Explanation: We can turn "1&(0|1)" into "1&(0&1)" by changing the '|' to a '&' using 1 operation.
      The new expression evaluates to 0.

  Example 2:
    Input: expression = "(0&0)&(0&0&0)"
    Output: 3
    Explanation: We can turn "(0&0)&(0&0&0)" into "(0|1)|(0&0&0)" using 3 operations.
      The new expression evaluates to 1.

  Example 3:
    Input: expression = "(0|(1|0&1))"
    Output: 1
    Explanation: We can turn "(0|(1|0&1))" into "(0|(0|0&1))" using 1 operation.
      The new expression evaluates to 0.

  Constraints:
    1. 1 <= expression.length <= 10^5
    2. expression only contains '1','0','&','|','(', and ')'
    3. All parentheses are properly matched.
    4. There will be no empty parentheses (i.e: "()" is not a substring of expression).

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/minimum-cost-to-change-the-final-value-of-expression
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Stack
// Dynamic Programming

func minOperationsToFlip(expression string) int {
	stk := IntStack{} // record previous '(', arr's lenght
	arr := [][2]int{} // [minOpTo 0, minOpTo 1]
	ops := []byte{}
	for i := range expression {
		op := expression[i]
		switch op {
		case '(':
			stk.Push(len(arr))
		case ')':
			start := stk.Pop()
			out := minOps(arr[start:], ops[start:])
			arr = arr[:start]
			ops = ops[:start]
			arr = append(arr, out)
		case '&', '|':
			ops = append(ops, op)
		case '0':
			arr = append(arr, [2]int{0, 1})
		case '1':
			arr = append(arr, [2]int{1, 0})
		}
	}

	out := minOps(arr, ops)

	if out[0] == 0 {
		return out[1]
	}
	return out[0]
}

func minOps(arr [][2]int, ops []byte) [2]int {
	n := len(arr)
	f := make([][2]int, n) // the min num of ops required to 0 or 1
	f[0] = arr[0]
	for i := 1; i < n; i++ {
		to0, to1 := arr[i][0], arr[i][1]
		if ops[i-1] == '&' {
			f[i][0] = Min(f[i-1][0]+to0, Min(f[i-1][1]+to0, f[i-1][0]+to1))   // 0 & 0 or 1 & 0 or 0 & 1
			f[i][1] = Min(f[i-1][1]+to1, Min(f[i-1][1]+to0, f[i-1][0]+to1)+1) // 1 & 1 or 1 | 0 or 0 | 1
		} else {
			f[i][0] = Min(f[i-1][0]+to0, Min(f[i-1][1]+to0, f[i-1][0]+to1)+1) // 0 | 0 or 0 & 1 or 1 & 0
			f[i][1] = Min(f[i-1][1]+to1, Min(f[i-1][1]+to0, f[i-1][0]+to1))   // 1 | 1 or 1 | 0 or 0 | 1
		}
	}
	return f[n-1]
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

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
