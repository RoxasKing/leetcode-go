package main

/*
  You are given a string word that consists of digits and lowercase English letters.

  You will replace every non-digit character with a space. For example, "a123bc34d8ef34" will become " 123  34 8  34". Notice that you are left with some integers that are separated by at least one space: "123", "34", "8", and "34".

  Return the number of different integers after performing the replacement operations on word.

  Two integers are considered different if their decimal representations without any leading zeros are different.

  Example 1:
    Input: word = "a123bc34d8ef34"
    Output: 3
    Explanation: The three different integers are "123", "34", and "8". Notice that "34" is only counted once.

  Example 2:
    Input: word = "leet1234code234"
    Output: 2

  Example 3:
    Input: word = "a1b01c001"
    Output: 1
    Explanation: The three integers "1", "01", and "001" all represent the same integer because
    the leading zeros are ignored when comparing their decimal values.

  Constraints:
    1. 1 <= word.length <= 1000
    2. word consists of digits and lowercase English letters.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/number-of-different-integers-in-a-string
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Two Pointers
func numDifferentIntegers(word string) int {
	mark := map[string]bool{"": true}
	cnt := 0
	l, r := 0, 0
	for ; r < len(word); r++ {
		if '0' <= word[r] && word[r] <= '9' {
			if l < r && word[l] == '0' {
				l++
			}
		} else {
			if !mark[word[l:r]] {
				cnt++
				mark[word[l:r]] = true
			}
			l = r + 1
		}
	}
	if !mark[word[l:r]] {
		cnt++
	}
	return cnt
}
