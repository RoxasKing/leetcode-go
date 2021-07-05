package main

func numJewelsInStones(J string, S string) int {
	dict := [128]bool{}
	for j := range J {
		dict[J[j]] = true
	}
	var count int
	for i := range S {
		if dict[S[i]] {
			count++
		}
	}
	return count
}
