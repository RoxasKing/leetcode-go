package main

import "sort"

// Tags:
// Math
// Sort

func orderlyQueue(s string, k int) string {
	n := len(s)
	out := []byte(s)
	if k == 1 {
		x := 0
		for i := 1; i < n; i++ {
			for j := 0; j < n; j++ {
				if out[(j+i)%n] > out[(j+x)%n] {
					break
				} else if out[(j+i)%n] < out[(j+x)%n] {
					x = i
					break
				}
			}
		}
		out = append(out[x:], out[:x]...)
	} else {
		sort.Slice(out, func(i, j int) bool { return out[i] < out[j] })
	}
	return string(out)
}
