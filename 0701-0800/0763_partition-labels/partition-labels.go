package main

/*
  A string S of lowercase English letters is given. We want to partition this string into as many parts as possible so that each letter appears in at most one part, and return a list of integers representing the size of these parts.

  Example 1:
    Input: S = "ababcbacadefegdehijhklij"
    Output: [9,7,8]
    Explanation:
      The partition is "ababcbaca", "defegde", "hijhklij".
      This is a partition so that each letter appears in at most one part.
      A partition like "ababcbacadefegde", "hijhklij" is incorrect, because it splits S into less parts.

  Note:
    1. S will have length in range [1, 500].
    2. S will consist of lowercase English letters ('a' to 'z') only.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/partition-labels
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Two Pointers + Greedy Algorithm
func partitionLabels(S string) []int {
	last := [26]int{}
	for i := range S {
		last[S[i]-'a'] = i
	}

	out := []int{}
	for l, r, i := 0, 0, 0; i < len(S); i++ {
		r = Max(r, last[S[i]-'a'])
		if i == r {
			out = append(out, r+1-l)
			l = r + 1
		}
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
