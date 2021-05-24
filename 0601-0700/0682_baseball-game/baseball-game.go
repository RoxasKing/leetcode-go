package main

import "strconv"

/*
  You are keeping score for a baseball game with strange rules. The game consists of several rounds, where the scores of past rounds may affect future rounds' scores.

  At the beginning of the game, you start with an empty record. You are given a list of strings ops, where ops[i] is the ith operation you must apply to the record and is one of the following:

    1. An integer x - Record a new score of x.
    2. "+" - Record a new score that is the sum of the previous two scores. It is guaranteed there will always be two previous scores.
    3. "D" - Record a new score that is double the previous score. It is guaranteed there will always be a previous score.
    4. "C" - Invalidate the previous score, removing it from the record. It is guaranteed there will always be a previous score.

  Return the sum of all the scores on the record.

  Example 1:
    Input: ops = ["5","2","C","D","+"]
    Output: 30
    Explanation:
      "5" - Add 5 to the record, record is now [5].
      "2" - Add 2 to the record, record is now [5, 2].
      "C" - Invalidate and remove the previous score, record is now [5].
      "D" - Add 2 * 5 = 10 to the record, record is now [5, 10].
      "+" - Add 5 + 10 = 15 to the record, record is now [5, 10, 15].
      The total sum is 5 + 10 + 15 = 30.

  Example 2:
    Input: ops = ["5","-2","4","C","D","9","+","+"]
    Output: 27
    Explanation:
      "5" - Add 5 to the record, record is now [5].
      "-2" - Add -2 to the record, record is now [5, -2].
      "4" - Add 4 to the record, record is now [5, -2, 4].
      "C" - Invalidate and remove the previous score, record is now [5, -2].
      "D" - Add 2 * -2 = -4 to the record, record is now [5, -2, -4].
      "9" - Add 9 to the record, record is now [5, -2, -4, 9].
      "+" - Add -4 + 9 = 5 to the record, record is now [5, -2, -4, 9, 5].
      "+" - Add 9 + 5 = 14 to the record, record is now [5, -2, -4, 9, 5, 14].
      The total sum is 5 + -2 + -4 + 9 + 5 + 14 = 27.

  Example 3:
    Input: ops = ["1"]
    Output: 1

  Constraints:
    1. 1 <= ops.length <= 1000
    2. ops[i] is "C", "D", "+", or a string representing an integer in the range [-3 * 10^4, 3 * 10^4].
    3. For operation "+", there will always be at least two previous scores on the record.
    4. For operations "C" and "D", there will always be at least one previous score on the record.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/baseball-game
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Stack

func calPoints(ops []string) int {
	s := IntStack{}
	for _, op := range ops {
		switch op {
		case "+":
			top := s.Len() - 1
			s.Push(s[top] + s[top-1])
		case "D":
			s.Push(s.Top() << 1)
		case "C":
			s.Pop()
		default:
			num, _ := strconv.Atoi(op)
			s.Push(num)
		}
	}

	out := 0
	for _, num := range s {
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
