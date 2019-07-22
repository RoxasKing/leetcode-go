package main

import "fmt"

func solveNQueens(n int) [][]string {
	out := [][]string{}
	// 用于记录第几行的第几列是否有值
	dicta := make(map[int]map[int]bool)
	dictb := make(map[int]bool)
	if n < 1 {
		return nil
	}
	// j 用于表示第一行处于第几列
	for i, j := 0, 0; i < n; i++ {
		// 用于存储解法
		strsa := []string{}
		str := "......."
		str = str[:i] + "Q" + str[i:]
		dictb[j] = true
		dicta[i] = dictb
		if strsb, ok := solve(dicta, j, n); ok {
			strsa = append(strsa, strsb...)
			// 将解法加到输出中
			out = append(out, strsa)
		}
	}
	return out
}

func solve(dict map[int]map[int]bool, j, n int) ([]string, bool) {
	j++
	if j >= n {
		return nil, false
	}
	// 遍历寻找
	for i := 0; i < n; i++ {
	}
}

func main() {
	fmt.Println(solveNQueens(4))
}
