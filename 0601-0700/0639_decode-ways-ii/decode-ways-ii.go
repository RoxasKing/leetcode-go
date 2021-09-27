package main

// Tags:
// Dynamic Programming

func numDecodings(s string) int {
	f0, f1, f2 := 0, 1, 0
	for i := range s {
		if s[i] == '*' {
			f2 += f1 * 9
		} else if s[i] != '0' {
			f2 += f1
		}
		if i > 0 && (s[i-1] == '1' || s[i-1] == '*') {
			f2 += f0
			if s[i] == '*' {
				f2 += f0 * 8
			}
		}
		if i > 0 && (s[i-1] == '2' || s[i-1] == '*') && (s[i] <= '6' || s[i] == '*') {
			f2 += f0
			if s[i] == '*' {
				f2 += f0 * 5
			}
		}
		f0, f1, f2 = f1, f2%mod, 0
	}
	return f1
}

const mod int = 1e9 + 7
