package main

import "sort"

/*
  Let the function f(s) be the frequency of the lexicographically smallest character in a non-empty string s. For example, if s = "dcce" then f(s) = 2 because the lexicographically smallest character is 'c', which has a frequency of 2.

  You are given an array of strings words and another array of query strings queries. For each query queries[i], count the number of words in words such that f(queries[i]) < f(W) for each W in words.

  Return an integer array answer, where each answer[i] is the answer to the ith query.

  Example 1:
    Input: queries = ["cbd"], words = ["zaaaz"]
    Output: [1]
    Explanation: On the first query we have f("cbd") = 1, f("zaaaz") = 3 so f("cbd") < f("zaaaz").

  Example 2:
    Input: queries = ["bbb","cc"], words = ["a","aa","aaa","aaaa"]
    Output: [1,2]
    Explanation: On the first query only f("bbb") < f("aaaa"). On the second query both f("aaa") and f("aaaa") are both > f("cc").

  Constraints:
    1. 1 <= queries.length <= 2000
    2. 1 <= words.length <= 2000
    3. 1 <= queries[i].length, words[i].length <= 10
    4. queries[i][j], words[i][j] consist of lowercase English letters.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/compare-strings-by-frequency-of-the-smallest-character
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Binary Search
func numSmallerByFrequency(queries []string, words []string) []int {
	m, n := len(queries), len(words)
	arr := make([]int, n)
	for i, word := range words {
		cnt := [26]int{}
		for j := range word {
			cnt[word[j]-'a']++
		}
		for j := 0; j < 26; j++ {
			if cnt[j] > 0 {
				arr[i] = cnt[j]
				break
			}
		}
	}
	sort.Ints(arr)

	qs := make([]int, m)
	for i, query := range queries {
		cnt := [26]int{}
		for j := range query {
			cnt[query[j]-'a']++
		}
		for j := 0; j < 26; j++ {
			if cnt[j] > 0 {
				qs[i] = cnt[j]
				break
			}
		}
	}

	out := make([]int, m)
	for i, q := range qs {
		j := sort.Search(n, func(j int) bool { return q < arr[j] })
		out[i] = n - j
	}
	return out
}
