package main

/*
  A string is a valid parentheses string (denoted VPS) if and only if it consists of "(" and ")" characters only, and:
    1. It is the empty string, or
    2. It can be written as AB (A concatenated with B), where A and B are VPS's, or
    3. It can be written as (A), where A is a VPS.

  We can similarly define the nesting depth depth(S) of any VPS S as follows:
    1. depth("") = 0
    2. depth(A + B) = max(depth(A), depth(B)), where A and B are VPS's
    3. depth("(" + A + ")") = 1 + depth(A), where A is a VPS.

  For example,  "", "()()", and "()(()())" are VPS's (with nesting depths 0, 1, and 2), and ")(" and "(()" are not VPS's.

  Given a VPS seq, split it into two disjoint subsequences A and B, such that A and B are VPS's (and A.length + B.length = seq.length).

  Now choose any such A and B such that max(depth(A), depth(B)) is the minimum possible value.

  Return an answer array (of length seq.length) that encodes such a choice of A and B:  answer[i] = 0 if seq[i] is part of A, else answer[i] = 1.  Note that even though multiple answers may exist, you may return any of them.

  Example 1:
    Input: seq = "(()())"
    Output: [0,1,1,1,1,0]

  Example 2:
    Input: seq = "()(())()"
    Output: [0,0,0,1,1,0,1,1]

  Constraints:
    1 <= seq.size <= 10000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/maximum-nesting-depth-of-two-valid-parentheses-strings
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func maxDepthexp1afterSplit(seq string) []int {
	out := make([]int, len(seq))
	depth := 0
	for i := range seq {
		if seq[i] == '(' {
			depth++
			out[i] = depth & 1
		}
		if seq[i] == ')' {
			out[i] = depth & 1
			depth--
		}
	}
	return out
}

func maxDepthexp1afterSplit2(seq string) []int {
	out := make([]int, len(seq))
	for i := range seq {
		if seq[i] == '(' {
			out[i] = 1 - i&1
		} else {
			out[i] = i & 1
		}
	}
	return out
}
