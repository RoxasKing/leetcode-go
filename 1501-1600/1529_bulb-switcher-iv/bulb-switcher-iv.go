package main

func minFlips(target string) int {
	out := int(target[0] - '0')
	for i := 1; i < len(target); i++ {
		if target[i] != target[i-1] {
			out++
		}
	}
	return out
}
