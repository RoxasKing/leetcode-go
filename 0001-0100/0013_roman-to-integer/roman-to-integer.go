package main

// Difficulty:
// Easy

// Tags:
// Math

func romanToInt(s string) int {
	h := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	o := h[s[0]]
	for i := 1; i < len(s); i++ {
		if o += h[s[i]]; h[s[i-1]] < h[s[i]] {
			o -= h[s[i-1]] * 2
		}
	}
	return o
}
