package main

import "sort"

/*
  Given a string array words, return the maximum value of length(word[i]) * length(word[j]) where the two words do not share common letters. If no such two words exist, return 0.

  Example 1:
    Input: words = ["abcw","baz","foo","bar","xtfn","abcdef"]
    Output: 16
    Explanation: The two words can be "abcw", "xtfn".

  Example 2:
    Input: words = ["a","ab","abc","d","cd","bcd","abcd"]
    Output: 4
    Explanation: The two words can be "ab", "cd".

  Example 3:
    Input: words = ["a","aa","aaa","aaaa"]
    Output: 0
    Explanation: No such pair of words.

  Constraints:
    1. 2 <= words.length <= 1000
    2. 1 <= words[i].length <= 1000
    3. words[i] consists only of lowercase English letters.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/maximum-product-of-word-lengths
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Bit Maniputlation
// Sort

func maxProduct(words []string) int {
	sort.Slice(words, func(i, j int) bool { return len(words[i]) > len(words[j]) })

	n := len(words)
	arr := make([]int, n)
	for j, w := range words {
		for i := range w {
			arr[j] |= 1 << int(w[i]-'a')
		}
	}

	out := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n && len(words[i])*len(words[j]) > out; j++ {
			if arr[i]^arr[j] == arr[i]+arr[j] {
				out = Max(out, len(words[i])*len(words[j]))
			}
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
