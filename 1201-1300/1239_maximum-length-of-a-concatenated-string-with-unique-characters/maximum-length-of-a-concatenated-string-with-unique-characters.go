package main

// Tags:
// Backtracking

func maxLength(arr []string) int {
	freq := make([]int, 26)
	var out int
	backtrack(arr, freq, 0, 0, &out)
	return out
}

func backtrack(arr []string, freq []int, i, cur int, out *int) {
	if i == len(arr) {
		return
	}
	ok := true
	for j := range arr[i] {
		freq[arr[i][j]-'a']++
		if freq[arr[i][j]-'a'] > 1 {
			ok = false
		}
	}
	if ok {
		cur += len(arr[i])
		*out = Max(*out, cur)
		backtrack(arr, freq, i+1, cur, out)
		cur -= len(arr[i])
	}
	for j := range arr[i] {
		freq[arr[i][j]-'a']--
	}
	backtrack(arr, freq, i+1, cur, out)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
