package main

/*
  Given a string, write a function to check if it is a permutation of a palindrome. A palindrome is a word or phrase that is the same forwards and backwards. A permutation is a rearrangement of letters. The palindrome does not need to be limited to just dictionary words.

  Example1:
    Input: "tactcoa"
    Output: true（permutations: "tacocat"、"atcocta", etc.）

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/palindrome-permutation-lcci
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Hash

func canPermutePalindrome(s string) bool {
	cnt := map[rune]int{}
	for _, ch := range s {
		cnt[ch]++
	}
	count := 0
	for _, v := range cnt {
		if v&1 == 1 {
			count++
		}
		if count > 1 {
			return false
		}
	}
	return true
}
