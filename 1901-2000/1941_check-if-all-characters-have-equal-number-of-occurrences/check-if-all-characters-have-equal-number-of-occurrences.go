package main

// Tags:
// Hash

func areOccurrencesEqual(s string) bool {
	freq := [26]int{}
	for i := range s {
		freq[s[i]-'a']++
	}
	arr := []int{}
	for i := 0; i < 26; i++ {
		if freq[i] > 0 {
			arr = append(arr, freq[i])
		}
	}
	for i := 1; i < len(arr); i++ {
		if arr[0] != arr[i] {
			return false
		}
	}
	return true
}
