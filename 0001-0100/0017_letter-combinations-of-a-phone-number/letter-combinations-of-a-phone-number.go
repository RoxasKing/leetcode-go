package main

// Difficulty:
// Medium

// Tags:
// Hash

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	h := []string{"", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	o := []string{""}
	for i := range digits {
		t := []string{}
		hi := h[digits[i]-'1']
		for _, p := range o {
			for j := range hi {
				t = append(t, p+hi[j:j+1])
			}
		}
		o = t
	}
	return o
}
