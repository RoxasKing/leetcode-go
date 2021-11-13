package main

// Difficulty:
// Hard

// Tags:
// Hash
// Bit Manipulation
// Backtracking

func findNumOfValidWords(words []string, puzzles []string) []int {
	mskCnt := map[int]int{}
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
		getSubMsks(puzzle, 1, 1<<(puzzle[0]-'a'), &subMsks)
		for _, msk := range subMsks {
			if cnt, ok := mskCnt[msk]; ok {
				out[i] += cnt
			}
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
