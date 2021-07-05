package main

// Tags:
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
