package main

/*
  Implement an algorithm to determine if a string has all unique characters. What if you cannot use additional data structures?

  Example 1:
    Input: s = "leetcode"
    Output: false

  Example 2:
    Input: s = "abc"
    Output: true

  Note:
    0 <= len(s) <= 100

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/is-unique-lcci
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Bit Manipulation

func isUnique(astr string) bool {
	cur := 0
	for i := 0; i < len(astr); i++ {
		if cur&(1<<int(astr[i]-'a')) > 0 {
			return false
		}
		cur |= 1 << int(astr[i]-'a')
	}
	return true
}
