package main

/*
  We define str = [s, n] as the string str which consists of the string s concatenated n times.
    For example, str == ["abc", 3] =="abcabcabc".
  We define that string s1 can be obtained from string s2 if we can remove some characters from s2 such that it becomes s1.
    For example, s1 = "abc" can be obtained from s2 = "abdbec" based on our definition by removing the bolded underlined characters.
  You are given two strings s1 and s2 and two integers n1 and n2. You have the two strings str1 = [s1, n1] and str2 = [s2, n2].

  Return the maximum integer m such that str = [str2, m] can be obtained from str1.

  Example 1:
    Input: s1 = "acb", n1 = 4, s2 = "ab", n2 = 2
    Output: 2

  Example 2:
    Input: s1 = "acb", n1 = 1, s2 = "acb", n2 = 1
    Output: 1

  Constraints:
    1. 1 <= s1.length, s2.length <= 100
    2. s1 and s2 consist of lowercase English letters.
    3. 1 <= n1, n2 <= 10^6

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/count-the-repetitions
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Important!

// Dynamic Programming
func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	if n1 == 0 {
		return 0
	}
	s1cnt, index, s2cnt := 0, 0, 0
	recall := make(map[int][]int)
	var preLoop, inLoop []int
	for {
		s1cnt++
		for i := range s1 {
			if s1[i] == s2[index] {
				index++
				if index == len(s2) {
					s2cnt, index = s2cnt+1, 0
				}
			}
		}
		if s1cnt == n1 {
			return s2cnt / n2
		}
		if _, ok := recall[index]; ok {
			preLoop = recall[index]
			inLoop = []int{s1cnt - recall[index][0], s2cnt - recall[index][1]}
			break
		} else {
			recall[index] = []int{s1cnt, s2cnt}
		}
	}
	ans := preLoop[1] + (n1-preLoop[0])/inLoop[0]*inLoop[1]
	rest := (n1 - preLoop[0]) % inLoop[0]
	for i := 0; i < rest; i++ {
		for j := range s1 {
			if s1[j] == s2[index] {
				index++
				if index == len(s2) {
					ans, index = ans+1, 0
				}
			}
		}
	}
	return ans / n2
}
