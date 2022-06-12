package main

// Difficulty:
// Hard

// Tags:
// Dynamic Programming

func countPalindromicSubsequences(s string) int {
	const mod int = 1e9 + 7
	n := len(s)
	f := make([][][]int, 4)
	for k := 0; k < 4; k++ {
		f[k] = make([][]int, n)
		for i := 0; i < n; i++ {
			f[k][i] = make([]int, n)
		}
	}
	f[s[0]-'a'][0][0] = 1
	for i := 1; i < n; i++ {
		f[s[i-1]-'a'][i-1][i] = 1
		f[s[i]-'a'][i-1][i] = 1
		if s[i-1] == s[i] {
			f[s[i]-'a'][i-1][i] = 2
		}
		f[s[i]-'a'][i][i] = 1
	}
	for l := n - 1; l >= 0; l-- {
		for r := l + 2; r < n; r++ {
			for k := 0; k < 4; k++ {
				c := byte(k) + 'a'
				if s[l] != c {
					f[k][l][r] = f[k][l+1][r]
				} else if s[r] != c {
					f[k][l][r] = f[k][l][r-1]
				} else {
					f[k][l][r] = 2
					for i := 0; i < 4; i++ {
						f[k][l][r] = (f[k][l][r] + f[i][l+1][r-1]) % mod
					}
				}
			}
		}
	}
	o := 0
	for k := 0; k < 4; k++ {
		o = (o + f[k][0][n-1]) % mod
	}
	return o
}
