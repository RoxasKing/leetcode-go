package main

// Tags:
// Backtracking
func splitIntoFibonacci(S string) []int {
	arr := []int{}
	if dfs(S, 0, &arr) {
		return arr
	}
	return []int{}
}

func dfs(S string, index int, arr *[]int) bool {
	if index == len(S) {
		return len(*arr) > 2
	}
	if S[index] == '0' {
		if len(*arr) >= 2 && 0 != (*arr)[len(*arr)-1]+(*arr)[len(*arr)-2] {
			return false
		}
		*arr = append(*arr, 0)
		if dfs(S, index+1, arr) {
			return true
		}
		*arr = (*arr)[:len(*arr)-1]
		return false
	}
	val := 0
	for i := index; i < len(S) && val < (1<<31-int(S[i]-'0'))/10; i++ {
		val = val*10 + int(S[i]-'0')
		if len(*arr) >= 2 && val < (*arr)[len(*arr)-1]+(*arr)[len(*arr)-2] {
			continue
		} else if len(*arr) >= 2 && val > (*arr)[len(*arr)-1]+(*arr)[len(*arr)-2] {
			break
		}
		*arr = append(*arr, val)
		if dfs(S, i+1, arr) {
			return true
		}
		*arr = (*arr)[:len(*arr)-1]
	}
	return false
}
