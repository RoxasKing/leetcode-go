package main

// Difficulty:
// Medium

func buildArray(target []int, n int) []string {
	o := []string{}
	for i, x := 0, 1; i < len(target) && x <= n; x++ {
		o = append(o, "Push")
		if target[i] != x {
			o = append(o, "Pop")
		} else {
			i++
		}
	}
	return o
}
