package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	a := len(word1)
	b := len(word2)

	save := make([][]int, a+1)
	for i := range save {
		save[i] = make([]int, b+1)
	}

	// word2 为空时，word1 变成 word2 的步骤
	for i := 1; i <= a; i++ {
		save[i][0] = i
	}

	// word1 为空时，word1 变成 word2 的步骤
	for j := 1; j <= b; j++ {
		save[0][j] = j
	}

	for i := 1; i <= a; i++ {
		for j := 1; j <= b; j++ {
			// 比较 减去 word1 最后一个字符和 减去 word2 最后一个字符后，、
			// 两种情况从 word1 变成 word2 谁的步骤最少，并加上减去这个字符的步骤
			save[i][j] = 1 + min(save[i-1][j], save[i][j-1])

			replace := 0
			// 若 word1 和 word2 当前坐标字符相同，则无需替换,总步骤和 save[i-1][j-1] 一致,
			// 若不同，则步骤数要再加上 1
			if word1[i-1] != word2[j-1] {
				replace = 1
			}

			// 比较最短替换步骤数，作为 i,j 坐标的值
			save[i][j] = min(save[i][j], save[i-1][j-1]+replace)
		}
	}

	return save[a][b]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	word1 := "intention"
	word2 := "execution"
	fmt.Println(minDistance(word1, word2))
}
