package main

// Difficulty:
// Medium

func minPartitions(n string) int {
	o := 0
	for _, c := range n {
		if x := int(c - '0'); o < x {
			o = x
		}
	}
	return o
}
