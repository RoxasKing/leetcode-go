package main

/*
  You are given two strings a and b that consist of lowercase letters. In one operation, you can change any character in a or b to any lowercase letter.

  Your goal is to satisfy one of the following three conditions:

    1. Every letter in a is strictly less than every letter in b in the alphabet.
    2. Every letter in b is strictly less than every letter in a in the alphabet.
    3. Both a and b consist of only one distinct letter.

  Return the minimum number of operations needed to achieve your goal.

  Example 1:
    Input: a = "aba", b = "caa"
    Output: 2
    Explanation: Consider the best way to make each condition true:
      1) Change b to "ccc" in 2 operations, then every letter in a is less than every letter in b.
      2) Change a to "bbb" and b to "aaa" in 3 operations, then every letter in b is less than every letter in a.
      3) Change a to "aaa" and b to "aaa" in 2 operations, then a and b consist of one distinct letter.
      The best way was done in 2 operations (either condition 1 or condition 3).

  Example 2:
    Input: a = "dabadd", b = "cda"
    Output: 3
    Explanation: The best way is to make condition 1 true by changing b to "eee".

  Constraints:
    1. 1 <= a.length, b.length <= 10^5
    2. a and b consist only of lowercase letters.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/change-minimum-characters-to-satisfy-one-of-three-conditions
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Greedy Algorithm

func minCharacters(a string, b string) int {
	freqA := [26]int{}
	freqB := [26]int{}

	for i := range a {
		freqA[a[i]-'a']++
	}
	for i := range b {
		freqB[b[i]-'a']++
	}

	out := len(a) + len(b)
	cntA, cntB := 0, 0
	for i := 0; i < 26; i++ {
		cntA += freqA[i]
		cntB += freqB[i]
		if i < 25 {
			out = Min(out, cntB+len(a)-cntA)
			out = Min(out, cntA+len(b)-cntB)
		}
		out = Min(out, len(a)+len(b)-(freqA[i]+freqB[i]))
	}

	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
