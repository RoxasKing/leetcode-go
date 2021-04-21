package main

/*
  With respect to a given puzzle string, a word is valid if both the following conditions are satisfied:
    1. word contains the first letter of puzzle.
    2. For each letter in word, that letter is in puzzle.
       For example, if the puzzle is "abcdefg", then valid words are "faced", "cabbage", and "baggage"; while invalid words are "beefed" (doesn't include "a") and "based" (includes "s" which isn't in the puzzle).
  Return an array answer, where answer[i] is the number of words in the given word list words that are valid with respect to the puzzle puzzles[i].

  Example :
    Input:
      words = ["aaaa","asas","able","ability","actt","actor","access"],
      puzzles = ["aboveyz","abrodyz","abslute","absoryz","actresz","gaswxyz"]
    Output: [1,1,3,2,4,0]
    Explanation:
      1 valid word for "aboveyz" : "aaaa"
      1 valid word for "abrodyz" : "aaaa"
      3 valid words for "abslute" : "aaaa", "asas", "able"
      2 valid words for "absoryz" : "aaaa", "asas"
      4 valid words for "actresz" : "aaaa", "asas", "actt", "access"
      There're no valid words for "gaswxyz" cause none of the words in the list contains letter 'g'.

  Constraints:
    1. 1 <= words.length <= 10^5
    2. 4 <= words[i].length <= 50
    3. 1 <= puzzles.length <= 10^4
    4. puzzles[i].length == 7
    5. words[i][j], puzzles[i][j] are English lowercase letters.
    6. Each puzzles[i] doesn't contain repeated characters.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/number-of-valid-words-for-each-puzzle
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Bit Manipulation + Hash + Backtracking
func findNumOfValidWords(words []string, puzzles []string) []int {
	mskCnt := make(map[int]int)
	for _, word := range words {
		msk := 0
		for j := range word {
			msk |= 1 << int(word[j]-'a')
		}
		mskCnt[msk]++
	}

	out := make([]int, len(puzzles))
	for i, puzzle := range puzzles {
		subMsks := []int{}
		msk := 1 << (puzzle[0] - 'a')
		getSubMsks(puzzle, 1, msk, &subMsks)
		for _, msk := range subMsks {
			out[i] += mskCnt[msk]
		}
	}
	return out
}

func getSubMsks(puzzle string, i int, msk int, subMsks *[]int) {
	if i == len(puzzle) {
		*subMsks = append(*subMsks, msk)
		return
	}

	getSubMsks(puzzle, i+1, msk, subMsks)
	getSubMsks(puzzle, i+1, msk|(1<<(puzzle[i]-'a')), subMsks)
}
