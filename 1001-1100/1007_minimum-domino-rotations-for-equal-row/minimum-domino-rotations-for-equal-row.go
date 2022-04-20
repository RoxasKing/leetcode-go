package main

// Difficulty:
// Medium

// Tags:
// Enumeration

func minDominoRotations(tops []int, bottoms []int) int {
	o := -1
	for k := 1; k <= 6; k++ {
		a, b := 0, 0
		ok := true
		for i := range tops {
			if tops[i] != k && bottoms[i] != k {
				ok = false
				break
			}
			if tops[i] != k {
				a++
			}
			if bottoms[i] != k {
				b++
			}
		}
		if !ok {
			continue
		}
		if o == -1 || o > a {
			o = a
		}
		if o == -1 || o > b {
			o = b
		}
	}
	return o
}
