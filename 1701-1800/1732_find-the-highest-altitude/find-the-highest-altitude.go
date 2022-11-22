package main

// Difficulty:
// Easy

// Tags:
// Prefix Sum

func largestAltitude(gain []int) int {
	o, v := 0, 0
	for _, x := range gain {
		if v += x; o < v {
			o = v
		}
	}
	return o
}
