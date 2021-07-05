package main

// Tags:
// Math

func romanToInt(s string) int {
	dict := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	out := 0
	for i := range s {
		out += dict[s[i]]
		if i > 0 && dict[s[i-1]] < dict[s[i]] {
			out -= 2 * dict[s[i-1]]
		}
	}
	return out
}
