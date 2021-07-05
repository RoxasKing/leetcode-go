package main

// Tags:
// Dynamic Programming
func getRow(rowIndex int) []int {
	out := make([]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		out[i] = 1
		for j := i - 1; j >= 1; j-- {
			out[j] += out[j-1]
		}
	}
	return out
}
