package main

// Difficulty:
// Medium

// Tags:
// Two Pointers

func reverseWords(s string) string {
	n := len(s)
	arr := []string{}
	for l, r := 0, 0; l < n; l = r {
		for ; l < n && s[l] == ' '; l++ {
		}
		for r = l; r < n && s[r] != ' '; r++ {
		}
		arr = append(arr, s[l:r])
	}
	chs := make([]byte, 0, n)
	for i := len(arr) - 1; i >= 0; i-- {
		if len(chs) > 0 {
			chs = append(chs, ' ')
		}
		chs = append(chs, []byte(arr[i])...)
	}
	return string(chs)
}
