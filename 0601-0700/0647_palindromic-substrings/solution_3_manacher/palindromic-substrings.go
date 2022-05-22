package main

import "strings"

// DiffiCulty:
// Medium

// Tags:
// Manacher

func countSubstrings(s string) int {
	ss := strings.Split(s, "")
	t := strings.Join(ss, "#")
	t = "$#" + t + "#!"
	n := len(t) - 1

	f := make([]int, n)
	iMax, rMax, out := 0, 0, 0
	for i := 1; i < n; i++ {
		if i < rMax {
			f[i] = Min(rMax-i+1, f[iMax-(i-iMax)])
		} else {
			f[i] = 1
		}

		for t[i+f[i]] == t[i-f[i]] {
			f[i]++
		}
		if i+f[i]-1 > rMax {
			iMax = i
			rMax = i + f[i] - 1
		}
		out += f[i] / 2
	}
	return out
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
