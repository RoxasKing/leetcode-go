package main

// Tags:
// Two Pointers
func longestBeautifulSubstring(word string) int {
	n := len(word)
	arr := make([]int, n)
	for i := range word {
		if word[i] == 'a' {
			arr[i] = 0
		} else if word[i] == 'e' {
			arr[i] = 1
		} else if word[i] == 'i' {
			arr[i] = 2
		} else if word[i] == 'o' {
			arr[i] = 3
		} else {
			arr[i] = 4
		}
	}

	out := 0

	for l, r := 0, 0; r < n; {
		if arr[l] != 0 {
			l, r = l+1, l+1
			continue
		}
		if r+1 < n && (arr[r+1] == arr[r] || arr[r+1] == arr[r]+1) {
			r++
			if arr[r] == 4 && r+1-l > out {
				out = r + 1 - l
			}
		} else {
			l, r = r+1, r+1
		}
	}

	return out
}
